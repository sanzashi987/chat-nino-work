package completion

import (
	"github.com/cza14h/chat-nino-work/model"
)

type DialogModel struct {
	model.BaseModel
	DialogConfig string         `gorm:"column:dialog_config;type:varchar(255)"`
	UserID       uint64         `gorm:"index"`
	Messages     []MessageModal `gorm:"foreignKey:DialogID"`
	IsDeleted    bool           `gorm:"index;default:false"`
}

type MessageModal struct {
	model.BaseModel
	DialogID  uint64 `gorm:"index"`
	IsDeleted bool   `gorm:"index;default:false"`
}

type DialogConfig struct {
	model.ChatConfig
	DialogTitle     string `json:"dialog_title"`
	UseGlobalConfig bool   `json:"use_global_config"`
}
