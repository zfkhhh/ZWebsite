package routers

import (
	"ZWebsite/pkg/setting"
	v1 "ZWebsite/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	mode := gin.ReleaseMode
	if setting.Setting.LogLevel == "debug" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", healthHandler)

	groupV1 := r.Group("/v1")
	v1.InitAccount(groupV1)

	return r
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}