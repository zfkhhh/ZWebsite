package admin

import (
	"ZWebsite/pkg/domain"
	"ZWebsite/pkg/e"
	"ZWebsite/pkg/logger"
	"ZWebsite/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAccount(group *gin.RouterGroup) {
	group.POST("/login", Login)
}

func Login(c *gin.Context) {
	ctx := e.GetContext(c)
	appG := e.Gin{C: c}

	loginReq := &domain.LoginReq{}

	if err := c.ShouldBind(loginReq);err != nil{
		logger.For(ctx).Errorf("invalid params [%v]", err)
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,false,map[string]string{
			"err_msg": err.Error(),
		})
		return
	}

	service := impl.NewService(ctx)

	uid,result, err := service.Login(loginReq.AccountName, loginReq.AccountPassword)
	if err != nil && !result{
		appG.Response(http.StatusBadRequest,e.ERROR_ACCOUNT_LOGIN,false,map[string]string{
			"err_msg": err.Error(),
		})
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,true,map[string]string{
		"uid": uid,
	})
	return
}
