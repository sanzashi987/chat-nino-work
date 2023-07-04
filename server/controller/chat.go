package controller

import (
	"net/http"

	"github.com/cza14h/chat-nino-work/dto"
	"github.com/cza14h/chat-nino-work/services"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	BaseController
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (c *ChatController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "",
	})

}

func (c *ChatController) Completion(ctx *gin.Context) {
	var requestBody dto.ChatMessageDto
	ctx.BindJSON(&requestBody)

	res, err := services.ReplyFromGPT(&requestBody)
	if err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, err.Error(), nil)
	}

}
