package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/repositories/twitter"
)

func New(twitterRepo twitter.Repository) *Impl {
	return &Impl{twitterRepo: twitterRepo}
}

func (service *Impl) Get(locale amqp.Language) ([]entities.WebhookTwitter, error) {
	return service.twitterRepo.Get(locale)
}

func (service *Impl) BatchUpdate(webhooks []entities.WebhookTwitter) error {
	return service.twitterRepo.BatchUpdate(webhooks)
}

func (service *Impl) BatchDelete(webhooks []entities.WebhookTwitter) error {
	return service.twitterRepo.BatchDelete(webhooks)
}
