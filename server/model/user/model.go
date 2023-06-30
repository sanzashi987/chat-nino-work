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

type UserConfig struct {
	Preference UserPreference   `json:"preference"`
	Chat       model.ChatConfig `json:"chat"`
}

type UserPreference struct {
	Avatar           string `json:"avatar"`
	SendKey          string `json:"send_key"`
	Theme            string `json:"theme"`
	Language         string `json:"language"`
	FontSize         string `json:"font_size"`
	SendPreviewubble bool   `json:"send_preview_bubble"`
	Mask             bool   `json:"mask"`
}
