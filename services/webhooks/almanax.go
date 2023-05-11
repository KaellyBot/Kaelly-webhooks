package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (service *Impl) GetAlmanaxWebhooks(locale amqp.Language) ([]entities.WebhookAlmanax, error) {
	return service.webhooksRepo.GetAlmanaxWebhooks(locale)
}

func (service *Impl) UpdateAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error {
	return service.webhooksRepo.UpdateAlmanaxWebhooks(webhooks)
}

func (service *Impl) DeleteAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error {
	return service.webhooksRepo.DeleteAlmanaxWebhooks(webhooks)
}
