package twitter

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(locale amqp.Language) ([]entities.WebhookTwitter, error) {
	var webhooks []entities.WebhookTwitter
	return webhooks, repo.db.GetDB().
		Where("locale = ?", locale).
		Find(webhooks).Error
}

func (repo *Impl) BatchUpdate(webhooks []entities.WebhookTwitter) error {
	// TODO use update instead of save to avoid weird things
	return repo.db.GetDB().Save(&webhooks).Error
}

func (repo *Impl) BatchDelete(webhooks []entities.WebhookTwitter) error {
	return repo.db.GetDB().Delete(&webhooks).Error
}
