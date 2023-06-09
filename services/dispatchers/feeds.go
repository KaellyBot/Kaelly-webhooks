package dispatchers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) dispatchFeed(correlationID string, feed *amqp.NewsRSSMessage, locale amqp.Language) {
	payload := mappers.MapFeed(feed, service.feedService, locale)
	webhooks, err := service.webhookService.GetFeedWebhooks(feed.Type, locale)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Cannot retrieve webhooks, ignoring tweet...")
		return
	}

	printedNumber := 0
	updatedWebhooks := make([]entities.WebhookFeed, 0)
	excludedWebhooks := make([]entities.WebhookFeed, 0)

	for _, webhook := range webhooks {
		err = service.discordService.PublishWebhook(webhook.WebhookID, webhook.WebhookToken, payload)
		if err != nil {
			log.Warn().Err(err).
				Str(constants.LogWebhookID, webhook.WebhookID).
				Msgf("Applying failure policy on feed webhook and continue...")

			if toKeep, updatedPolicy := service.applyFailurePolicy(webhook.WebhookPolicy); toKeep {
				webhook.WebhookPolicy = updatedPolicy
				updatedWebhooks = append(updatedWebhooks, webhook)
			} else {
				excludedWebhooks = append(excludedWebhooks, webhook)
			}
		} else {
			webhook.WebhookPolicy = service.applySuccessPolicy(webhook.WebhookPolicy)
			updatedWebhooks = append(updatedWebhooks, webhook)
			printedNumber++
		}
	}

	log.Info().
		Str(constants.LogCorrelationID, correlationID).
		Int(constants.LogPrintNumber, printedNumber).
		Int(constants.LogWebhookNumber, len(webhooks)).
		Msgf("Feed has been published!")

	service.updateFeedWebhooks(updatedWebhooks)
	service.deleteFeedWebhooks(excludedWebhooks)
}

func (service *Impl) updateFeedWebhooks(webhooks []entities.WebhookFeed) {
	err := service.webhookService.UpdateFeedWebhooks(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot update feed webhooks, ignoring them for this time")
	}
}

func (service *Impl) deleteFeedWebhooks(webhooks []entities.WebhookFeed) {
	err := service.webhookService.DeleteFeedWebhooks(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot remove unreachable feed webhooks, ignoring them for this time")
	}
}
