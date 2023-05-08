package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/repositories/almanax"
)

type Service interface {
	Get(locale amqp.Language) ([]entities.WebhookAlmanax, error)
	BatchUpdate(webhooks []entities.WebhookAlmanax) error
	BatchDelete(webhooks []entities.WebhookAlmanax) error
}

type Impl struct {
	almanaxRepo almanax.Repository
}
