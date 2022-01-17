package dao

import (
	"ZWebsite/pkg/constant"
	"ZWebsite/pkg/domain"
	"ZWebsite/pkg/setting"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"k8s.io/klog/v2"
)

var (
	DB          *gorm.DB
	dbInfo      = &domain.DBInfo{}
	allModels   = []interface{}{&RequestLog{}}
)

func Init() {


	if dbInfo.User = setting.Setting.DBUser; dbInfo.User == "" {
		klog.Fatalf("get db user failed")
	}
	if dbInfo.Password = setting.Setting.DBPassword; dbInfo.Password == "" {
		klog.Fatalf("get db password failed")
	}
	if dbInfo.Host = setting.Setting.DBHost; dbInfo.Host == "" {
		klog.Fatalf("get db host failed")
	}
	if dbInfo.Port = setting.Setting.DBPort; dbInfo.Port == "" {
		klog.Fatalf("get db port failed")
	}
	if dbInfo.Name = setting.Setting.DBName; dbInfo.Name == "" {
		klog.Fatalf("get db name failed")
	}

	dbInfo.Path = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbInfo.User, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Name)
	klog.Infof("database path is [%v]", dbInfo.Path)

	var err error
	DB, err = gorm.Open(mysql.Open(dbInfo.Path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		klog.Fatalf("fail to setup db [%v]", err)
	}

	DB.Logger = DB.Logger.LogMode(logger.LogLevel(int(logger.Info)))

	sqlDB, err := DB.DB()
	if err != nil {
		klog.Fatalf("failed to get sql DB [%v]", err)
	}
	sqlDB.SetMaxIdleConns(constant.DBMaxIdleConn)
	sqlDB.SetMaxOpenConns(constant.DBMaxOpenConn)
	sqlDB.SetConnMaxLifetime(constant.DBConnMaxLifeTime)

	// InnoDB
	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(allModels...); err != nil {
		klog.Fatalf("fail to automigrate database [%v]", err)
	}

	klog.Info("init DB connection success")
}