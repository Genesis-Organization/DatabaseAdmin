package app

import (
	. "GenesisDAT/src/config"
	controllers "GenesisDAT/src/controllers"

	"github.com/gin-gonic/gin"
)

func AppInit() {
	App := gin.Default()

	App.GET("/", controllers.ServeIndex)
	App.Run(Config.Port)
}
