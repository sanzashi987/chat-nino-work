package completion

import (
	"github.com/cza14h/chat-nino-work/model"
)

type DialogModel struct {
	model.BaseModel
	DialogConfig string         `gorm:"column:dialog_config;type:varchar(255)"`
	UserID       uint           `gorm:"index"`
	Messages     []MessageModal `gorm:"foreignKey:DialogID"`
	IsDeleted    bool           `gorm:"index;default:false"`
}

type MessageModal struct {
	model.BaseModel
	DialogID  uint `gorm:"index"`
	IsDeleted bool `gorm:"index;default:false"`
}

func GetPagingDialogsByUserID(userId uint, pageSize int, pageIndex int) (dialogs *[]DialogModel) {
	model.DBRef.Where("user_id = ? AND is_deleted = ?", userId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(dialogs)
	return dialogs
}

func GetPagingMessagsByDialogID(dialogId uint, pageSize int, pageIndex int) (messages *[]MessageModal) {
	model.DBRef.Where("dialog_id = ? AND is_deleted = ?", dialogId, false).Offset(pageIndex * pageSize).Limit(pageSize).Find(messages)
	return messages
}

func MarkDialogIsDeleted(dialogId uint) {
	dialog := &DialogModel{}
	model.DBRef.First(dialog, dialogId).Update("is_deleted", true)
}
