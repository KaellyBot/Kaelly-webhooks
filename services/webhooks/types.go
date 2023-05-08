package webhooks

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/services/almanax"
	"github.com/kaellybot/kaelly-webhooks/services/discord"
	"github.com/kaellybot/kaelly-webhooks/services/feeds"
	"github.com/kaellybot/kaelly-webhooks/services/twitter"
)

const (
	webhookQueueName  = "webhook"
	webhookRoutingkey = "news.*"

	maxRetry = 10
	delta    = 2 * time.Hour
)

type Service interface {
	Consume() error
}

type Impl struct {
	broker         amqp.MessageBroker
	almanaxService almanax.Service
	feedService    feeds.Service
	twitterService twitter.Service
	discordService discord.Service
}
