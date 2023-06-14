package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	BaseController
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (*ChatController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "",
	})

}
