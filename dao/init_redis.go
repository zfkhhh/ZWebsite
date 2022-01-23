package dao

import (
	"ZWebsite/pkg/setting"
	"fmt"
	"github.com/go-redis/redis"
	"k8s.io/klog/v2"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", setting.Setting.RedisHost, setting.Setting.RedisPort),
		Password: setting.Setting.RedisPassword,
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		klog.Fatalf("ping redis failed , err = [%v]",err)
	}

	klog.Fatalf("Init Redis Success")
}

