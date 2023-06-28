package user

import (
	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/services/dto"
	"github.com/cza14h/chat-nino-work/utils"
	"gorm.io/gorm"
)

func UpdateUserConfig(userId uint, payload *dto.UpdateUserConfigDto) (err error) {
	updateUser := UserModel{
		PreferenceConfig: payload.PreferenceConfig,
		ChatConfig:       payload.PreferenceConfig,
	}
	err = model.DBRef.Where("id = ? ", userId).Updates(&updateUser).Error
	return
}

func ReadByUsername(username string) (user *UserModel, err error) {
	user = &UserModel{}
	err = model.DBRef.Where("username = ?", username).First(&user).Error
	return
}

func ReadByUserID(userId uint) (user *UserModel, err error) {
	user = &UserModel{}
	err = model.DBRef.First(user, userId).Error
	return user, err
}

func (user *UserModel) ComparPassword(_password string) bool {
	return utils.CheckHash(_password, user.Password)
}

// Gorm hook
func (user *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	if !utils.IsHashed(user.Password) {
		user.Password = utils.MakeHash(user.Password)
	}
	return
}
