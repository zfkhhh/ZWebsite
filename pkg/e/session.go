package e

import (
	"ZWebsite/pkg/constant"
	"ZWebsite/pkg/logger"
	"ZWebsite/pkg/setting"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"strconv"
)

func SetUpSession() *gin.HandlerFunc {
	redisAddr := fmt.Sprintf("%s:%s", setting.Setting.RedisHost, setting.Setting.RedisPort)
	redisConNum, _ := strconv.Atoi(constant.RedisConNum)
	store, err := redis.NewStore(
		redisConNum,
		constant.RedisNetwork,
		redisAddr,
		setting.Setting.RedisPassword,
		[]byte(constant.RedisSessionSecret),
	)
	if err != nil {
		klog.Fatalf("failed to connect redis")
	}
	handlerFunc := sessions.Sessions(setting.Setting.ServiceName, store)
	return &handlerFunc
}

func SetSession(c *gin.Context, uid string) error {
	session := sessions.Default(c)
	// session alive 60 min
	option := sessions.Options{MaxAge: 3600}
	session.Options(option)
	session.Set(uid , uid)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetSession(c *gin.Context,uid string) bool {
	session := sessions.Default(c)
	userId := session.Get(uid)
	logger.For(c).Infof("user [%v] is login ", userId)
	if userId != nil {
		return true
	} else {
		return false
	}
}
