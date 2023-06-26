package controller

import "github.com/gin-gonic/gin"

type BaseController struct {
}

type BasePageSizePayload struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

func (*BaseController) RespondJson(c *gin.Context, code int, errMsg string, data interface{}) {
	c.JSON(code, gin.H{
		"code":   code,
		"errMsg": errMsg,
		"data":   data,
	})
}

func (*BaseController) AbortJson(c *gin.Context, code int, errMsg string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"errMsg": errMsg,
		"data":   data,
	})
	panic(errMsg)
}
