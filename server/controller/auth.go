package controller

import "github.com/gin-gonic/gin"

type AuthController struct {
	BaseController
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (auth *AuthController) Login(ctx *gin.Context) {
	var loginPayload = LoginPayload{}
	ctx.BindJSON(&loginPayload)

}
