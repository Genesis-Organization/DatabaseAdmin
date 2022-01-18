package router

import (
	Controllers "GenesisDAT/src/controllers"

	gin "github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	R := gin.Default()
	R.LoadHTMLGlob("src/web/**/*.html")

	R.Static("/libs", "src/web/libs")
	R.Static("/assets", "src/web/assets")

	R.SetTrustedProxies([]string{"192.168.1.2"})

	R.GET("/", Controllers.ServeIndex)

	return R
}
