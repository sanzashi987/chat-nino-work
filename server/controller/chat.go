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

func (c *ChatController) Chat(ctx *gin.Context) {
	var requestBody dto.ChatMessageDto
	ctx.BindJSON(&requestBody)

	res, messageId, err := services.ReplyFromGPT(ctx, &requestBody)
	if err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	c.RespondJson(ctx, http.StatusOK, "", dto.ResponseChatMessageDto{
		Content:       res,
		UserMessageId: messageId,
	})

}

func (c *ChatController) ReChat(ctx *gin.Context) {
	var requestBody dto.RequestReChatMessageDto
	ctx.BindJSON(&requestBody)

}
