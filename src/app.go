package src

import (
	Config "GenesisDAT/src/config"
	Router "GenesisDAT/src/router"

	gin "github.com/gin-gonic/gin"
)

func App() {
	if Config.Mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	App := Router.CreateServer()

	App.Run(Config.Port)
}

func Init() {
	Logs()

	App()

}
