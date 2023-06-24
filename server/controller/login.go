package controller

import (
	"github.com/cza14h/chat-nino-work/config"
	"github.com/cza14h/chat-nino-work/model/user"
	"gorm.io/gorm"

	"net/http"

	authPkg "github.com/cza14h/chat-nino-work/pkg/auth"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (auth *LoginController) Login(ctx *gin.Context) {
	var loginPayload = LoginPayload{}
	err := ctx.BindJSON(&loginPayload)
	if err != nil {
		auth.AbortJson(ctx, http.StatusBadRequest, "Invalid login parameters", nil)
	}

	user, err := user.GetByUsername(loginPayload.Username)
	if err != nil && err == gorm.ErrRecordNotFound {
		auth.AbortJson(ctx, http.StatusUnauthorized, "User not found", nil)
	}

	if ok := user.ComparPassword(loginPayload.Username); !ok {
		auth.AbortJson(ctx, http.StatusUnauthorized, "Password not match", nil)
	}

	token, err := authPkg.GenerateToken(user.Username, uint(user.ID))
	if err != nil {
		auth.AbortJson(ctx, http.StatusInternalServerError, "Fail to generate login token", nil)
	}

	ctx.SetCookie(config.JwtTokenHeader, token, int(config.JwtCookieExpiry.Seconds()), "/", "*", false, true)

	ctx.Redirect(http.StatusFound, "/")
}
