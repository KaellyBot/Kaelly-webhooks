package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(feedTypeID string, locale amqp.Language) ([]entities.WebhookFeed, error) {
	var webhooks []entities.WebhookFeed
	return webhooks, repo.db.GetDB().
		Where("feed_type_id = ? AND locale = ?",
			feedTypeID, locale).
		Find(webhooks).Error
}

func (repo *Impl) BatchUpdate(webhooks []entities.WebhookFeed) error {
	// TODO use update instead of save to avoid weird things
	return repo.db.GetDB().Save(&webhooks).Error
}

func (repo *Impl) BatchDelete(webhooks []entities.WebhookFeed) error {
	return repo.db.GetDB().Delete(&webhooks).Error
}
