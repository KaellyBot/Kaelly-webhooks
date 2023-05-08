package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

type Repository interface {
	Get(locale amqp.Language) ([]entities.WebhookTwitter, error)
	BatchUpdate(webhooks []entities.WebhookTwitter) error
	BatchDelete(webhooks []entities.WebhookTwitter) error
}

type Impl struct {
	db databases.MySQLConnection
}
