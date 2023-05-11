package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (service *Impl) GetFeedWebhooks(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error) {
	return service.webhooksRepo.GetFeedWebhooks(feedTypeID, locale)
}

func (service *Impl) UpdateFeedWebhooks(webhooks []entities.WebhookFeed) error {
	return service.webhooksRepo.UpdateFeedWebhooks(webhooks)
}

func (service *Impl) DeleteFeedWebhooks(webhooks []entities.WebhookFeed) error {
	return service.webhooksRepo.DeleteFeedWebhooks(webhooks)
}
