package e

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	RequestId   string      `json:"requestId"`
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, success bool, data interface{}) {

	g.C.JSON(httpCode, Response{

		Code:    errCode,
		Data:    data,
	})
	return
}
