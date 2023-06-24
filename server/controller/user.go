package controller

import (
	"net/http"

	"github.com/cza14h/chat-nino-work/config"
	"github.com/cza14h/chat-nino-work/model/user"

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
	userId, ok := ctx.MustGet(config.JwtUserIDContextKey).(uint)
	if !ok {
		auth.AbortJson(ctx, http.StatusInternalServerError, "Fail to get token from context", nil)
	}
	return userId
}

func (auth *AuthController) Info(ctx *gin.Context) {
	userId := mustGetUserID(ctx, auth)
	userModel, _ := user.GetByUserID(userId)

}

type QueryMessages struct {
	DialogID uint `json:"dialog_id"`
}

func (auth *AuthController) Messages(ctx *gin.Context) {

}

type QueryPagingDialogs struct {
	BasePageSize
}

func (auth *AuthController) PagingDialogs(ctx *gin.Context) {
	userId := mustGetUserID(ctx, auth)

}

type QueryPagingMessages struct {
	QueryMessages
	BasePageSize
}

func (auth *AuthController) PagingMessages(ctx *gin.Context) {

}
