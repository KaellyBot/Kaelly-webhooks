package dispatchers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) dispatchTweet(correlationID string, tweet *amqp.NewsTwitterMessage, locale amqp.Language) {
	payload := mappers.MapTweet(tweet)
	webhooks, err := service.webhookService.GetTwitterWebhooks(locale)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Cannot retrieve webhooks, ignoring tweet...")
		return
	}

	printedNumber := 0
	updatedWebhooks := make([]entities.WebhookTwitter, 0)
	excludedWebhooks := make([]entities.WebhookTwitter, 0)

	for _, webhook := range webhooks {
		err = service.discordService.PublishWebhook(webhook.WebhookID, webhook.WebhookToken, payload)
		if err != nil {
			log.Warn().Err(err).
				Str(constants.LogWebhookID, webhook.WebhookID).
				Msgf("Applying failure policy on twitter webhook and continue...")

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
		Msgf("Tweet has been published!")

	service.updateTwitterWebhooks(updatedWebhooks)
	service.deleteTwitterWebhooks(excludedWebhooks)
}

func (service *Impl) updateTwitterWebhooks(webhooks []entities.WebhookTwitter) {
	err := service.webhookService.UpdateTwitterWebhooks(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot update twitter webhooks, ignoring them for this time")
	}
}

func (service *Impl) deleteTwitterWebhooks(webhooks []entities.WebhookTwitter) {
	err := service.webhookService.DeleteTwitterWebhooks(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot remove unreachable twitter webhooks, ignoring them for this time")
	}
}
