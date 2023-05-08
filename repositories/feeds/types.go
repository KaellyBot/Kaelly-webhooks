package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

type Repository interface {
	Get(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error)
	BatchUpdate(webhooks []entities.WebhookFeed) error
	BatchDelete(webhooks []entities.WebhookFeed) error
}

type Impl struct {
	db databases.MySQLConnection
}
