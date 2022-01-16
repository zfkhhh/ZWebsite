package e

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success     bool         `json:"success"`
	AppMessage  *AppError    `json:"appMessage"`
	Data        interface{}  `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, success bool, data interface{}) {

	g.C.JSON(httpCode, Response{
		AppMessage: NewErr(errCode),
		Success:    success,
		Data:       data,
	})
	return
}
