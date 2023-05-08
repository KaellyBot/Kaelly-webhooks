package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/repositories/feeds"
)

func New(feedsRepo feeds.Repository) *Impl {
	return &Impl{feedsRepo: feedsRepo}
}

func (service *Impl) Get(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error) {
	return service.feedsRepo.Get(feedTypeID, locale)
}

func (service *Impl) BatchUpdate(webhooks []entities.WebhookFeed) error {
	return service.feedsRepo.BatchUpdate(webhooks)
}

func (service *Impl) BatchDelete(webhooks []entities.WebhookFeed) error {
	return service.feedsRepo.BatchDelete(webhooks)
}
