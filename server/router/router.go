package router

import (
	"github.com/cza14h/chat-nino-work/controller"
	"github.com/cza14h/chat-nino-work/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = controller.NewChatController()
var authController = controller.NewAuthController()
var loginController = controller.NewLoginController()

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())

	// frontend files
	router.GET("/", chatController.Index)

	// public apis
	login := router.Group("login")
	{
		login.POST("/login", loginController.Login)
	}

	// auth required apis
	chat := router.Group("chat").Use(middlewares.Jwt())
	{
		chat.POST("")
	}

	auth := router.Group("user").Use(middlewares.Jwt())
	{
		auth.POST("/info", authController.Info)

	}

}
