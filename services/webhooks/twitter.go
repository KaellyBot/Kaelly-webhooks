package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (service *Impl) GetTwitterWebhooks(locale amqp.Language) ([]entities.WebhookTwitter, error) {
	return service.webhooksRepo.GetTwitterWebhooks(locale)
}

func (service *Impl) UpdateTwitterWebhooks(webhooks []entities.WebhookTwitter) error {
	return service.webhooksRepo.UpdateTwitterWebhooks(webhooks)
}

func (service *Impl) DeleteTwitterWebhooks(webhooks []entities.WebhookTwitter) error {
	return service.webhooksRepo.DeleteTwitterWebhooks(webhooks)
}
