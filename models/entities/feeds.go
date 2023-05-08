package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookFeed struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	FeedTypeID   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	WebhookPolicy
}
