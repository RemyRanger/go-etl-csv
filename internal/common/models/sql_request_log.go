package models

import (
	"time"
)

// RequestLog : database RequestLog entity model
type RequestLog struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        int64   `gorm:"primary_key;type:int"`
	Payload   *string `gorm:"column:payload;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (rl *RequestLog) TableName() string {
	return "request_logs"
}
