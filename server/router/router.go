package router

import (
	"github.com/cza14h/chat-nino-work/controller"
	"github.com/cza14h/chat-nino-work/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = controller.NewChatController()
var authController = controller.NewChatController()

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())
	router.GET("/", chatController.Index)

	chat := router.Group("chat")
	{
		chat.POST("")
	}

	auth := router.Group("auth")
	{
		auth.POST("")
	}
}
