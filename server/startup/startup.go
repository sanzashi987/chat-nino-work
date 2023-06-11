package startup

import (
	"strconv"

	"github.com/cza14h/chat-nino-work/config"
)

func StartApp() {

	appConfig := config.LoadConfig()

	serverPortInString := strconv.Itoa(appConfig.Port)

}

func startStaticServer() {

}
