package setting

import (
	"ZWebsite/pkg/constant"
	"ZWebsite/pkg/utils"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Settings struct {
	ServiceName string
	ServicePort string

	LogLevel string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDBIndex  string
}

var Setting = &Settings{}

func init() {
	if !utils.Exists(constant.AppConfigPath) {
		log.Fatalf("app config not exist")
	}

	v := viper.New()
	v.SetConfigFile(constant.AppConfigPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("fail to read config")
	}

	// service name
	if Setting.ServiceName = os.Getenv("servicename"); Setting.ServiceName == "" {
		Setting.ServiceName = v.GetString("service")
	}
	// service port
	if Setting.ServicePort = os.Getenv("port"); Setting.ServicePort == "" {
		Setting.ServicePort = v.GetString("port")
	}
	// log level
	if Setting.LogLevel = os.Getenv("loglevel"); Setting.LogLevel == "" {
		log.Fatalf("fail to get loglevel")
	}
	// mysql settings
	if Setting.DBHost = os.Getenv("DBHOST"); Setting.DBHost == "" {
		log.Fatalf("failed to get DBHOST")
	}
	if Setting.DBPort = os.Getenv("DBPORT"); Setting.DBPort == "" {
		log.Fatalf("failed to get DBPORT")
	}
	if Setting.DBUser = os.Getenv("DBUSER"); Setting.DBUser == "" {
		log.Fatalf("failed to get DBUSER")
	}
	if Setting.DBPassword = os.Getenv("DBPW"); Setting.DBPassword == "" {
		log.Fatalf("failed to get DBPW")
	}
	if Setting.DBName = os.Getenv("DBNAME"); Setting.DBName == "" {
		log.Fatalf("failed to get DBNAME")
	}
	if Setting.RedisHost = os.Getenv("REDISHOST"); Setting.RedisHost == "" {
		log.Fatalf("failed to get RedisHost")
	}
	if Setting.RedisPort = os.Getenv("REDISPORT"); Setting.RedisPort == "" {
		log.Fatalf("failed to get RedisPort")
	}
	// Redis password allow ""
	Setting.RedisPassword = os.Getenv("REDISPW")
	if Setting.RedisDBIndex = os.Getenv("REDISDBINDEX"); Setting.RedisDBIndex == "" {
		// default index = 0
		Setting.RedisDBIndex = constant.RedisDBIndex
		log.Printf("Redis default DBIndex")
	}
}
