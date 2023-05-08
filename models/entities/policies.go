package entities

import "time"

type WebhookPolicy struct {
	RetryNumber int64 `gorm:"default:0"`
	UpdatedAt   time.Time
}
