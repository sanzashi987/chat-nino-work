package controller

import (
	"fmt"
	"net/http"

	"github.com/cza14h/chat-nino-work/consts"
	"github.com/cza14h/chat-nino-work/dto"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/cza14h/chat-nino-work/services"

	// authPkg "github.com/cza14h/chat-nino-work/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func mustGetUserID(ctx *gin.Context, auth *AuthController) uint {
	userId, ok := ctx.MustGet(consts.JwtUserIDContextKey).(uint)
	if !ok {
		auth.AbortJson(ctx, http.StatusInternalServerError, "Fail to get token from context", nil)
		panic("Fail to get token from context")
	}
	return userId
}

func (auth *AuthController) GetUserInfo(ctx *gin.Context) {
	requestUserInfo := dto.RequestDialogsDto{}
	userId := mustGetUserID(ctx, auth)
	ctx.BindJSON(&requestUserInfo)

	resDto, err := services.GetUserInfo(userId, &requestUserInfo)

	if err != nil {
		auth.AbortJson(ctx, http.StatusInternalServerError,
			fmt.Sprintf("fail to fetch user, err: %s", err.Error()), nil)
		return
	}

	auth.RespondJson(ctx, http.StatusOK, "", &resDto)

}

func (auth *AuthController) Test(ctx *gin.Context) {
	requestUserInfo := dto.RequestDialogsDto{}

	if err := ctx.BindJSON(&requestUserInfo); err != nil {
		auth.AbortJson(ctx, 400, "", nil)
		return
	}

	fmt.Println(requestUserInfo)
	resDto := dto.ResponseUserInfoDto{
		UpdateUserConfigDto: dto.UpdateUserConfigDto{
			PreferenceConfig: "233",
		},
	}

	auth.RespondJson(ctx, http.StatusOK, "", &resDto)

}

func (auth *AuthController) Messages(ctx *gin.Context) {
	var requestBody = dto.RequestMessagesDto{
		RequestPagingDto: dto.RequestPagingDto{
			PageSize:  10,
			PageIndex: 0,
		}, // default pagination config
	}
	ctx.BindJSON(&requestBody)
	userId := mustGetUserID(ctx, auth)

	legal, dialog := completion.IsDialogBelongsToUser(userId, requestBody.DialogID)
	if !legal {
		auth.AbortJson(ctx, http.StatusBadRequest, "dialog not belongs to the user", nil)
		return
	}

	messages, _ := completion.ReadPagingMessagsByDialogID(
		requestBody.DialogID,
		requestBody.PageSize,
		requestBody.PageIndex,
	)

	auth.RespondJson(ctx, http.StatusOK, "", gin.H{
		"messages":     messages,
		"pageIndex":    requestBody.PageIndex,
		"pageSize":     requestBody.PageSize,
		"messageCount": dialog.MessageCount,
	})

}

func (auth *AuthController) PagingDialogs(ctx *gin.Context) {
	// userId := mustGetUserID(ctx, auth)

}

func (auth *AuthController) PagingMessages(ctx *gin.Context) {

}
