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

	// version

	v1 := router.Group("V1")

	// public apis
	login := v1.Group("login")
	{
		login.POST("/login", loginController.Login)
	}

	// auth required apis
	chat := v1.Group("chat").Use(middlewares.Jwt())
	{
		chat.POST("/completion", chatController.Completion)
	}

	auth := v1.Group("user").Use(middlewares.Jwt())
	{
		auth.POST("/info", authController.GetUserInfo)
	}

	test := v1.Group("test")
	{
		test.POST("/test", authController.Test)
	}

}
