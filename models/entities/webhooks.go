package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookPolicy struct {
	RetryNumber int64 `gorm:"default:0"`
	UpdatedAt   time.Time
}

type WebhookAlmanax struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	WebhookPolicy
}

type WebhookFeed struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	FeedTypeID   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	WebhookPolicy
}

type WebhookTwitter struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	WebhookPolicy
}
