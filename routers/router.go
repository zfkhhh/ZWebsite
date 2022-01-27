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
	r.Use()

	r.GET("/", healthHandler)

	groupAdmin := r.Group("/v1/admin")
	v1.InitAdmin(groupAdmin)

	groupWeb := r.Group("/v1/website")
	v1.InitWebSite(groupWeb)

	return r
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}
