package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/repositories/almanax"
)

func New(almanaxRepo almanax.Repository) *Impl {
	return &Impl{almanaxRepo: almanaxRepo}
}

func (service *Impl) Get(locale amqp.Language) ([]entities.WebhookAlmanax, error) {
	return service.almanaxRepo.Get(locale)
}

func (service *Impl) BatchUpdate(webhooks []entities.WebhookAlmanax) error {
	return service.almanaxRepo.BatchUpdate(webhooks)
}

func (service *Impl) BatchDelete(webhooks []entities.WebhookAlmanax) error {
	return service.almanaxRepo.BatchDelete(webhooks)
}
