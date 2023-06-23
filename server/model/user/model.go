package user

import (
	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/cza14h/chat-nino-work/utils"
	"gorm.io/gorm"
)

type UserModel struct {
	model.BaseModel
	UserConfig string                   `gorm:"column:user_config;type:varchar(255)"`
	UserName   string                   `gorm:"column:user_name;type:varchar(255);unique"`
	Password   string                   `gorm:"column:password;type:varchar(255)"`
	Dialogs    []completion.DialogModel `gorm:"foreignKey:UserID"`
}

// Gorm hook
func (user *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	if !utils.IsHashed(user.Password) {
		user.Password = utils.MakeHash(user.Password)
	}
	return
}

// func (UserModel) TableName() string {
// 	return "users"
// }

func GetByUsername(userName string) (*UserModel, error) {
	user := UserModel{}
	err := model.DBRef.Where("user_name = ?", userName).First(&user).Error
	return &user, err
}

func (user *UserModel) ComparPassword(_password string) bool {
	return utils.CheckHash(_password, user.Password)
}
