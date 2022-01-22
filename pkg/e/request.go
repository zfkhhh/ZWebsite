package e

import (
	"ZWebsite/pkg/constant"
	"ZWebsite/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
)

func GetContext(c *gin.Context) context.Context {
	requestId := utils.NewUUID()
	ctx := context.WithValue(c.Request.Context(), constant.RequestKey, requestId)
	return ctx
}
