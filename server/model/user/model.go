package user

import (
	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/utils"
	"gorm.io/gorm"
)

type UserModel struct {
	model.BaseModel
	UserConfig string `gorm:"column:user_config;type:varchar(255)"`
	UserName   string `gorm:"column:user_name;type:varchar(255);unique"`
	Password   string `gorm:"column:password;type:varchar(255)"`
}

func (user *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	if !utils.IsHashed(user.Password) {
		user.Password = utils.MakeHash(user.Password)
	}
	return
}
