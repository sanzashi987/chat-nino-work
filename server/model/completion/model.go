package completion

import (
	"github.com/cza14h/chat-nino-work/model"
)

type DialogModel struct {
	model.BaseModel
	IsDelete     bool   `gorm:"colume:is_deleted"`
	DialogConfig string `gorm:"column:dialog_config;type:varchar(255)"`
	UserID       uint
	Messages     []MessageModal `gorm:"foreignKey:DialogID"`
}

type MessageModal struct {
	model.BaseModel
	DialogID uint
}
