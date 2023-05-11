package feeds

import (
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
)

type Repository interface {
	GetFeedTypes() ([]entities.FeedType, error)
}

type Impl struct {
	db databases.MySQLConnection
}
