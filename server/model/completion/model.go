package completion

import (
	"github.com/cza14h/chat-nino-work/model"
)

type DialogModel struct {
	model.BaseModel
	DialogConfig string         `gorm:"column:dialog_config;type:varchar(255)"`
	UserID       uint64         `gorm:"index"`
	Messages     []MessageModel `gorm:"foreignKey:DialogID"`
	IsDeleted    bool           `gorm:"index;default:false;type:boolean"`
	MessageCount int            `gorm:"column:message_count"`
}

func (u DialogModel) TableName() string {
	return "dialogs"
}

type MessageModel struct {
	model.BaseModel
	Content  string `gorm:"column:content;type:varchar(255)"`
	DialogID uint64 `gorm:"index"`
	/**
	 * `0` indicates the current record is user's message
	 * otherwise indicates the message replies an user's message,
	 * and the value is the id of the message
	 */
	ReplyTo   uint64 `gorm:"column:reply_to;default:0"`
	IsDeleted bool   `gorm:"index;default:false;type:boolean"`
}

func (u MessageModel) TableName() string {
	return "messages"
}
