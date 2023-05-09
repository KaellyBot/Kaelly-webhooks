package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) dispatchAlmanax(correlationID string, almanax *amqp.NewsRSSMessage, locale amqp.Language) {
	payload := mappers.MapAlmanax(almanax)
	webhooks, err := service.almanaxService.Get(locale)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, correlationID).
			Msgf("Cannot retrieve webhooks, ignoring tweet...")
		return
	}

	printedNumber := 0
	updatedWebhooks := make([]entities.WebhookAlmanax, 0)
	excludedWebhooks := make([]entities.WebhookAlmanax, 0)

	for _, webhook := range webhooks {
		err = service.discordService.PublishWebhook(webhook.WebhookID, webhook.WebhookToken, payload)
		if err != nil {
			log.Warn().Err(err).
				Str(constants.LogWebhookID, webhook.WebhookID).
				Msgf("Applying failure policy on almanax webhook and continue...")

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
		Msgf("Almanax has been published!")

	service.updateAlmanaxWebhooks(updatedWebhooks)
	service.deleteAlmanaxWebhooks(excludedWebhooks)
}

func (service *Impl) updateAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) {
	err := service.almanaxService.BatchUpdate(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot update almanax webhooks, ignoring them for this time")
	}
}

func (service *Impl) deleteAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) {
	err := service.almanaxService.BatchDelete(webhooks)
	if err != nil {
		log.Error().Err(err).
			Msgf("Cannot remove unreachable almanax webhooks, ignoring them for this time")
	}
}
