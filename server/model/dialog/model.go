package dialog

import "github.com/cza14h/chat-nino-work/model"

type DialogModel struct {
	model.BaseModel
	DialogConfig string `gorm:"column:dialog_config;type:varchar(255)"`
}
