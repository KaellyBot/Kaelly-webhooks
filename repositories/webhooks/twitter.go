package webhooks

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (repo *Impl) GetTwitterWebhooks(locale amqp.Language) ([]entities.WebhookTwitter, error) {
	var webhooks []entities.WebhookTwitter
	return webhooks, repo.db.GetDB().
		Where("locale = ?", locale).
		Find(&webhooks).Error
}

func (repo *Impl) UpdateTwitterWebhooks(webhooks []entities.WebhookTwitter) error {
	var err error
	for _, wh := range webhooks {
		webhook := wh
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Updates(webhook).Error)
	}
	return err
}

func (repo *Impl) DeleteTwitterWebhooks(webhooks []entities.WebhookTwitter) error {
	var err error
	for _, wh := range webhooks {
		webhook := wh
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Delete(webhook).Error)
	}
	return err
}
