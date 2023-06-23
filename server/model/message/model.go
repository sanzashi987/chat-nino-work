package message

import (
	"github.com/cza14h/chat-nino-work/model"
)

type MessageModal struct {
	model.BaseModel
}

func (MessageModal) TableName() string {
	return "messages"
}
