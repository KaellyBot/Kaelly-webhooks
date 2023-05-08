package almanax

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

type Repository interface {
	Get(locale amqp.Language) ([]entities.WebhookAlmanax, error)
	BatchUpdate(webhooks []entities.WebhookAlmanax) error
	BatchDelete(webhooks []entities.WebhookAlmanax) error
}

type Impl struct {
	db databases.MySQLConnection
}
