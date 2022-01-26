package main

import (
	"ZWebsite/dao"
	"ZWebsite/pkg/logger"
	"ZWebsite/pkg/setting"
	"ZWebsite/routers"
	"context"
	"fmt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
)

func main() {
	ctx := context.TODO()

	logger.SetUp()
	dao.InitMysql()
	dao.InitRedis()

	logger.For(ctx).Info("Starting server ......")

	r := routers.InitRouter()

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(fmt.Sprintf(":%s", setting.Setting.ServicePort))
}
