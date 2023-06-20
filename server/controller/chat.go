package controller

import (
	"net/http"

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

var supportModels = map[string]bool{
	openai.GPT4:              true,
	openai.GPT40314:          true,
	openai.GPT432K:           true,
	openai.GPT432K0314:       true,
	openai.GPT3Dot5Turbo0301: true,
	openai.GPT3Dot5Turbo:     true,
}

func (c *ChatController) Completion(ctx *gin.Context) {
	var request openai.ChatCompletionRequest
	ctx.BindJSON(&request)

	gptConfig := openai.DefaultConfig("")

	client := openai.NewClientWithConfig(gptConfig)

	if ok := supportModels[request.Model]; ok {
		response, err := client.CreateChatCompletion(ctx, request)
		if err != nil {
			c.RespondJson(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}
		c.RespondJson(ctx, http.StatusOK, "", gin.H{
			"response": response,
		})
	} else {
		c.RespondJson(ctx, http.StatusBadRequest, "supported model", gin.H{})
	}
}
