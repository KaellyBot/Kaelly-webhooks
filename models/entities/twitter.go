package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookTwitter struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	WebhookPolicy
}
