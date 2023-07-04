package user

import (
	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/model/completion"
)

type UserModel struct {
	model.BaseModel
	PreferenceConfig string                   `gorm:"column:preference_config;type:varchar(255)"`
	ChatConfig       string                   `gorm:"column:chat_config;type:varchar(255)"`
	Username         string                   `gorm:"column:username;type:varchar(255);unique"`
	Password         string                   `gorm:"column:password;type:varchar(255)"`
	Dialogs          []completion.DialogModel `gorm:"foreignKey:UserID"`
	DialogCount      int                      `gorm:"column:dialog_count"`
}

func (u UserModel) TableName() string {
	return "users"
}

