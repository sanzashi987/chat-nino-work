package controller

import (
	"net/http"

	"github.com/cza14h/chat-nino-work/consts"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
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
	var request openai.ChatCompletionRequest
	ctx.BindJSON(&request)

	gptConfig := openai.DefaultConfig("")

	client := openai.NewClientWithConfig(gptConfig)

	if ok := consts.SupportModels[request.Model]; ok {
		response, err := client.CreateChatCompletion(ctx, request)
		if err != nil {
			c.AbortJson(ctx, http.StatusBadRequest, err.Error(), nil)
		}
		c.RespondJson(ctx, http.StatusOK, "", gin.H{
			"response": response,
		})
	} else {
		c.AbortJson(ctx, http.StatusBadRequest, "unsupported model", gin.H{})
	}
}
