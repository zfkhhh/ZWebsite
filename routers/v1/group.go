package v1

import (
	"ZWebsite/routers/v1/admin"
	"ZWebsite/routers/v1/website"
	"github.com/gin-gonic/gin"
)

func InitAdmin(group *gin.RouterGroup) {
	admin.InitAccount(group)
}

func InitWebSite(group *gin.RouterGroup) {
	website.InitBlog(group)
}
