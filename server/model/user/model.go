package user

import (
	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/cza14h/chat-nino-work/utils"
	"gorm.io/gorm"
)

type UserModel struct {
	model.BaseModel
	UserConfig   string                   `gorm:"column:user_config;type:varchar(255)"`
	Username     string                   `gorm:"column:username;type:varchar(255);unique"`
	Password     string                   `gorm:"column:password;type:varchar(255)"`
	Dialogs      []completion.DialogModel `gorm:"foreignKey:UserID"`
	DialogCounts int                      `gorm:"column:dialog_counts;type:int"`
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

func GetByUsername(username string) (*UserModel, error) {
	user := UserModel{}
	err := model.DBRef.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetByUserID(userId uint) (user *UserModel, err error) {
	user = &UserModel{}
	err = model.DBRef.Where("ID = ?", userId).First(user).Error
	return user, err
}

func (user *UserModel) ComparPassword(_password string) bool {
	return utils.CheckHash(_password, user.Password)
}
