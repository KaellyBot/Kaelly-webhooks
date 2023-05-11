package webhooks

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (repo *Impl) GetFeedWebhooks(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error) {
	var webhooks []entities.WebhookFeed
	return webhooks, repo.db.GetDB().
		Where("feed_type_id = ? AND locale = ?",
			feedTypeID, locale).
		Find(&webhooks).Error
}

func (repo *Impl) UpdateFeedWebhooks(webhooks []entities.WebhookFeed) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Updates(webhook).Error)
	}
	return err
}

func (repo *Impl) DeleteFeedWebhooks(webhooks []entities.WebhookFeed) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Delete(webhook).Error)
	}
	return err
}
