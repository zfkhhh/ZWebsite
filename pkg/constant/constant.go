package constant

import "time"

const (
	AppConfigPath = "./conf/app.yaml"
	// RequestKey  request key name
	RequestKey = "request_uuid"

	AccountNameMinLength    = 1
	AccountNameMaxLength    = 20
	AccountPasswordMinLenth = 8
	AccountPasswordMaxLenth = 20
)

// db settings
const (
	DBMaxIdleConn     = 10
	DBMaxOpenConn     = 1024
	DBConnMaxLifeTime = 60 * time.Second

	RedisDBIndex = "0"
)

// log settings
const (
	// Log save dir
	LogDir = "/tmp/ZWebsite/log"
	// 日志保留最大时长环境变量
	LogRotateDaysEnvKey = "LOG_ROTATA_DAYS"
	// 日志保留最大时长默认值
	DefaultLogRotateDays = 30
	// 日志保留最大时长最小天数
	MinLogRotateDays = 7
)
