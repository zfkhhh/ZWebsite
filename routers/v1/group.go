package v1

import (
	"ZWebsite/routers/v1/admin"
	"github.com/gin-gonic/gin"
)

func InitAccount(group *gin.RouterGroup) {
	admin.InitAccount(group)
}