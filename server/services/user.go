package services

import (
	// "github.com/cza14h/chat-nino-work/model/user"
	"github.com/cza14h/chat-nino-work/services/dto"
)

func GetUserInfo(userId uint, payload *dto.RequestDialogsDto) (res dto.ResponseUserInfoDto, err error) {
	// always start paging from 0
	payload.PageIndex = 0
	// userModel, err := user.ReadByUserID(userId)
	// if err != nil {
	// 	return res, err
	// }

	return res, nil
}
