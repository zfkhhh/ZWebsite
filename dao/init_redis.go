package dao

import (
	"ZWebsite/pkg/setting"
	"fmt"
	"github.com/go-redis/redis"
	"k8s.io/klog/v2"
	"strconv"
)

var RedisClient *redis.Client

func InitRedis() {
	redisDBIndx ,_ := strconv.Atoi(setting.Setting.RedisDBIndex)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", setting.Setting.RedisHost, setting.Setting.RedisPort),
		Password: setting.Setting.RedisPassword,
		DB:       redisDBIndx,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		klog.Fatalf("ping redis failed , err = [%v]",err)
	}

	klog.Info("Init Redis Success")
}

