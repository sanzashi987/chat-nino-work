package router

import (
	"github.com/cza14h/chat-nino-work/controller"
	"github.com/cza14h/chat-nino-work/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = controller.NewChatController()
var authController = controller.NewAuthController()

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())

	chat := router.Group("chat").Use(middlewares.Jwt())
	{
		chat.POST("")
	}

	auth := router.Group("auth")
	{
		auth.POST("login", authController.Login)
	}

	router.GET("/", chatController.Index)
}
