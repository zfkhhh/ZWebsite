package dao

import (
	"ZWebsite/pkg/logger"
	"context"
	"gorm.io/gorm"
)

type RequestLog struct {
	gorm.Model

	Source    string
	ReqUrl    string
	ReqMethod string
	ReqBody   string `gorm:"type:text"`
	RespCode  string
	RespBody  string `gorm:"type:text"`
	SourceIP  string
}

func AddRequestLog(ctx context.Context, source, reqUrl, reqMethod, reqBody, respCode, respBody, sourceIP string) error {
	requestLog := &RequestLog{
		Source:    source,
		ReqUrl:    reqUrl,
		ReqMethod: reqMethod,
		ReqBody:   reqBody,
		RespCode:  respCode,
		RespBody:  respBody,
		SourceIP:  sourceIP,
	}
	if len(requestLog.ReqBody) > 10000 || len(requestLog.RespBody) > 10000 {
		logger.For(ctx).Info("big data, skip logging")
		return nil
	}
	if err := DB.Create(requestLog).Error; err != nil {
		return err
	}

	return nil
}
