package bootstrap

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/cza14h/chat-nino-work/config"
	"github.com/cza14h/chat-nino-work/model"
	iRouter "github.com/cza14h/chat-nino-work/router"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var once sync.Once

func registerRoutes() {
	once.Do(func() {
		router = gin.Default()
		iRouter.RegisterRoutes(router)
	})
}

func ConnectDb() {

	model.ConnectDB()

}

func setTemplateDir() {
	router.LoadHTMLFiles("static/*")
}

func startStaticServer() {
	router.GET("/assets/:filename", func(ctx *gin.Context) {
		filename := ctx.Param("filename")
		ext := filepath.Ext(filename)
		if ext == "js" {
			ctx.Header("Content-Type", "text/javascript")
		} else {
			ctx.Header("Content-Type", mime.TypeByExtension(ext))

		}

		path := "static/assets/" + filename

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.File(path)
	})
}

func StartApp() {
	registerRoutes()
	setTemplateDir()
	startStaticServer()
	appConfig := config.LoadConfig()
	serverPortInString := strconv.Itoa(appConfig.Port)
	router.Run(":" + serverPortInString)

}
