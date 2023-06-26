package completion

import (
	"time"

	"github.com/cza14h/chat-nino-work/config"
	"github.com/cza14h/chat-nino-work/model"
)

func CreateDialogFromConfig(userId uint64) {

	var dialog = &DialogModel{
		BaseModel: model.BaseModel{
			ID:        uint64(config.SnowflakeNode.Generate()),
			CreatedAt: time.Now(),
		},
		UserID: userId,
	}

}

func ReadPagingDialogsByUserID(userId uint, pageSize int, pageIndex int) (dialogs *[]DialogModel) {
	model.DBRef.Where("user_id = ? AND is_deleted = ?", userId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(dialogs)
	return
}

func ReadPagingMessagsByDialogID(dialogId uint, pageSize int, pageIndex int) (messages *[]MessageModal) {
	model.DBRef.Where("dialog_id = ? AND is_deleted = ?", dialogId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(messages)
	return
}

func IsDialogBelongsToUser(userId uint, dialogId uint) bool {
	var dialog = &DialogModel{}
	model.DBRef.First(dialog, dialogId)
	return dialog.UserID == uint64(userId)
}

func MarkDialogIsDeleted(dialogId uint) {
	dialog := &DialogModel{}
	model.DBRef.First(dialog, dialogId).Update("is_deleted", true)
}
