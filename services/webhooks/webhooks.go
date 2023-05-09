package webhooks

import (
	"context"
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/services/almanax"
	"github.com/kaellybot/kaelly-webhooks/services/discord"
	"github.com/kaellybot/kaelly-webhooks/services/feeds"
	"github.com/kaellybot/kaelly-webhooks/services/twitter"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBroker, almanaxService almanax.Service,
	feedService feeds.Service, twitterService twitter.Service,
	discordService discord.Service) *Impl {
	return &Impl{
		broker:         broker,
		almanaxService: almanaxService,
		feedService:    feedService,
		twitterService: twitterService,
		discordService: discordService,
	}
}

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeNews,
		RoutingKey: webhookRoutingkey,
		Queue:      webhookQueueName,
	}
}

func (service *Impl) Consume() error {
	log.Info().Msgf("Consuming news messages...")
	return service.broker.Consume(webhookQueueName, service.consume)
}

func (service *Impl) consume(_ context.Context,
	message *amqp.RabbitMQMessage, correlationID string) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	// case amqp.RabbitMQMessage_NEWS_ALMANAX:
	// TODO
	case amqp.RabbitMQMessage_NEWS_RSS:
		service.dispatchFeed(correlationID, message.NewsRSSMessage, message.Language)
	case amqp.RabbitMQMessage_NEWS_TWITTER:
		service.dispatchTweet(correlationID, message.NewsTwitterMessage, message.Language)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Type not recognized, request ignored")
	}
}

func (service *Impl) applySuccessPolicy(webhookPolicy entities.WebhookPolicy) entities.WebhookPolicy {
	if webhookPolicy.RetryNumber != 0 {
		webhookPolicy.RetryNumber = 0
		// TODO does not work
	}
	return webhookPolicy
}

func (service *Impl) applyFailurePolicy(webhookPolicy entities.WebhookPolicy) (bool, entities.WebhookPolicy) {
	if webhookPolicy.UpdatedAt.Add(delta).Before(time.Now()) {
		webhookPolicy.RetryNumber++
		if webhookPolicy.RetryNumber >= maxRetry {
			return false, webhookPolicy
		}
	}
	return true, webhookPolicy
}
