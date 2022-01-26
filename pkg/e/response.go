package e

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success    bool        `json:"success"`
	AppMessage *AppError   `json:"appMessage"`
	Data       interface{} `json:"data"`
}

func (g *Gin) Session() *Session {
	var session Session
	cookie, ok := g.C.Get("_session")
	if !ok {
		return nil
	}

	session = cookie.(Session)

	session.Lock = &sync.Mutex{}
	return &session
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
