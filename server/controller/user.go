package controller

import (
	"fmt"
	"net/http"

	"github.com/cza14h/chat-nino-work/consts"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/cza14h/chat-nino-work/services"
	"github.com/cza14h/chat-nino-work/services/dto"

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

type QueryMessagesPayload struct {
	BasePageSizePayload
	DialogID uint `json:"dialog_id" binding:"required"`
}

func (auth *AuthController) Messages(ctx *gin.Context) {
	var query = QueryMessagesPayload{
		BasePageSizePayload: BasePageSizePayload{
			PageIndex: 0,
			PageSize:  10,
		},
	}
	ctx.BindJSON(&query)
	userId := mustGetUserID(ctx, auth)

	legal, dialog := completion.IsDialogBelongsToUser(userId, query.DialogID)
	if !legal {
		auth.AbortJson(ctx, http.StatusBadRequest, "dialog not belongs to the user", nil)
		return
	}

	messages, _ := completion.ReadPagingMessagsByDialogID(
		query.DialogID,
		query.PageSize,
		query.PageIndex,
	)

	auth.RespondJson(ctx, http.StatusOK, "", gin.H{
		"messages":     messages,
		"pageIndex":    query.PageIndex,
		"pageSize":     query.PageSize,
		"messageCount": dialog.MessageCount,
	})

}

type QueryPagingDialogs struct {
	BasePageSizePayload
}

func (auth *AuthController) PagingDialogs(ctx *gin.Context) {
	// userId := mustGetUserID(ctx, auth)

}

type QueryPagingMessages struct {
	QueryMessagesPayload
	BasePageSizePayload
}

func (auth *AuthController) PagingMessages(ctx *gin.Context) {

}
