package webhooks

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
)

func (repo *Impl) GetAlmanaxWebhooks(locale amqp.Language) ([]entities.WebhookAlmanax, error) {
	var webhooks []entities.WebhookAlmanax
	return webhooks, repo.db.GetDB().
		Where("locale = ?", locale).
		Find(&webhooks).Error
}

func (repo *Impl) UpdateAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Updates(webhook).Error)
	}
	return err
}

func (repo *Impl) DeleteAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Delete(webhook).Error)
	}
	return err
}
