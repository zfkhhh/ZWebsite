package admin

import (
	"github.com/gin-gonic/gin"
)

func InitAccount(group *gin.RouterGroup) {
	group.POST("/login", Login)
}

func Login(c *gin.Context) {
}
