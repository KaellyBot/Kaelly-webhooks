package feeds

import (
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	repository "github.com/kaellybot/kaelly-webhooks/repositories/feeds"
)

type Service interface {
	FindFeedTypeByID(id string) (entities.FeedType, bool)
}

type Impl struct {
	feedTypes  map[string]entities.FeedType
	repository repository.Repository
}
