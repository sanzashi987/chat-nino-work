package middlewares

import (
	"net/http"

	"github.com/cza14h/chat-nino-work/config"
	"github.com/cza14h/chat-nino-work/controller"
	"github.com/cza14h/chat-nino-work/pkg/auth"
	"github.com/gin-gonic/gin"
)

var base = controller.BaseController{}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.ParseToken(c)
		if err != nil {
			base.AbortJson(c, http.StatusUnauthorized, err.Error(), nil)
		}
		c.Set(config.JwtContextKey, claims)
		c.Next()
	}
}
