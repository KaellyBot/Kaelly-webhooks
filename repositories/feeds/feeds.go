package feeds

import (
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetFeedTypes() ([]entities.FeedType, error) {
	var feedTypes []entities.FeedType
	response := repo.db.GetDB().Model(&entities.FeedType{}).Preload("Labels").Find(&feedTypes)
	return feedTypes, response.Error
}
