package constant

import "time"

const (
	AppConfigPath = "./conf/app.yaml"
	// RequestKey  request key name
	RequestKey = "request_uuid"
)

const (
	DBMaxIdleConn     = 10
	DBMaxOpenConn     = 1024
	DBConnMaxLifeTime = 60 * time.Second
)