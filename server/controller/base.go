package controller

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (*BaseController) RespondJson(c *gin.Context, code int, errMsg string, data interface{}) {
	c.JSON(code, gin.H{
		"code":   code,
		"errMsg": errMsg,
		"data":   data,
	})
}
