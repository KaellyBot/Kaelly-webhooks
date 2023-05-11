package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

type Repository interface {
	GetAlmanaxWebhooks(locale amqp.Language) ([]entities.WebhookAlmanax, error)
	GetFeedWebhooks(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error)
	GetTwitterWebhooks(locale amqp.Language) ([]entities.WebhookTwitter, error)
	UpdateAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error
	UpdateFeedWebhooks(webhooks []entities.WebhookFeed) error
	UpdateTwitterWebhooks(webhooks []entities.WebhookTwitter) error
	DeleteAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error
	DeleteFeedWebhooks(webhooks []entities.WebhookFeed) error
	DeleteTwitterWebhooks(webhooks []entities.WebhookTwitter) error
}

type Impl struct {
	db databases.MySQLConnection
}
