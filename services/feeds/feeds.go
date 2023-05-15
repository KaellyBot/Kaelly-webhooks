package feeds

import (
	"github.com/kaellybot/kaelly-webhooks/models/entities"
	repository "github.com/kaellybot/kaelly-webhooks/repositories/feeds"
)

func New(repository repository.Repository) (*Impl, error) {
	feedTypeEntities, err := repository.GetFeedTypes()
	if err != nil {
		return nil, err
	}

	feedTypes := make(map[string]entities.FeedType)
	for _, feedType := range feedTypeEntities {
		feedTypes[feedType.ID] = feedType
	}

	return &Impl{
		feedTypes:  feedTypes,
		repository: repository,
	}, nil
}

func (service *Impl) FindFeedTypeByID(id string) (entities.FeedType, bool) {
	feedType, found := service.feedTypes[id]
	return feedType, found
}
