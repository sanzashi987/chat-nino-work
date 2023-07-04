package completion

import (
	"github.com/cza14h/chat-nino-work/model"
)

func CreateDialogFromConfig(userId uint64) {

	var dialog = DialogModel{
		BaseModel: *model.CreateBaseModel(),
		UserID:    userId,
	}

	model.DBRef.Create(&dialog)
}

func ReadPagingDialogsByUserID(userId uint, pageIndex int, pageSize int) (dialogs *[]DialogModel, err error) {
	err = model.DBRef.Where("user_id = ? AND is_deleted = ?", userId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(dialogs).Error
	return
}

func ReadPagingMessagsByDialogID(dialogId uint, pageIndex int, pageSize int) (messages *[]MessageModel, err error) {
	err = model.DBRef.Where("dialog_id = ? AND is_deleted = ?", dialogId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(messages).Error
	return
}

func IsDialogBelongsToUser(userId uint, dialogId uint) (bool, *DialogModel) {
	var dialog = DialogModel{}
	model.DBRef.First(&dialog, dialogId)
	return dialog.UserID == uint64(userId), &dialog
}

func DeleteDialog(dialogId uint) (err error) {
	dialog := DialogModel{}
	err = model.DBRef.First(&dialog, dialogId).Update("is_deleted", true).Error
	return
}

func CreateMessagePair(replayContent, userContent string, dialogID uint64) (uint64, uint64, error) {
	var userMessage = MessageModel{
		BaseModel: *model.CreateBaseModel(),
		DialogID:  dialogID,
		Content:   userContent,
	}
	var assistantMessage = MessageModel{
		BaseModel: *model.CreateBaseModel(),
		ReplyTo:   userMessage.ID,
		DialogID:  dialogID,
		Content:   replayContent,
	}

	records := []*MessageModel{&userMessage, &assistantMessage}

	err := model.DBRef.Create(records).Error
	return userMessage.ID, assistantMessage.ID, err
}
