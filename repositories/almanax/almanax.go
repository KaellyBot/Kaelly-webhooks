package almanax

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(locale amqp.Language) ([]entities.WebhookAlmanax, error) {
	var webhooks []entities.WebhookAlmanax
	return webhooks, repo.db.GetDB().
		Where("locale = ?", locale).
		Find(&webhooks).Error
}

func (repo *Impl) BatchUpdate(webhooks []entities.WebhookAlmanax) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Updates(webhook).Error)
	}
	return err
}

func (repo *Impl) BatchDelete(webhooks []entities.WebhookAlmanax) error {
	var err error
	for _, webhook := range webhooks {
		err = errors.Join(err, repo.db.GetDB().Model(&webhook).Delete(webhook).Error)
	}
	return err
}
