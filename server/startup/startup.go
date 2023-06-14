package startup

import (
	"mime"
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

func connectDb() {
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
		ctx.File("static/assets" + filename)
	})
}

func StartApp() {
	registerRoutes()
	connectDb()
	startStaticServer()
	appConfig := config.LoadConfig()
	serverPortInString := strconv.Itoa(appConfig.Port)
	router.Run(":" + serverPortInString)

}
