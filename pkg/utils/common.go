package utils

import (
	uuid "github.com/satori/go.uuid"
)

// NewUUID 生成长id
func NewUUID() string {
	return uuid.NewV4().String()
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsValidPage(page int) bool {
	if page < 1 {
		return false
	}
	return true
}

func IsValidPageSize(pageSize int) bool {
	if pageSize > 0 && pageSize <= 100 {
		return true
	}
	return false
}
