package utils

import (
	uuid "github.com/satori/go.uuid"
)
// NewUUID 生成长id
func NewUUID() string {
	return uuid.NewV4().String()
}
