package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	feedsRepo "github.com/kaellybot/kaelly-webhooks/repositories/feeds"
	webhooksRepo "github.com/kaellybot/kaelly-webhooks/repositories/webhooks"
	"github.com/kaellybot/kaelly-webhooks/services/discord"
	"github.com/kaellybot/kaelly-webhooks/services/dispatchers"
	"github.com/kaellybot/kaelly-webhooks/services/feeds"
	"github.com/kaellybot/kaelly-webhooks/services/webhooks"
	"github.com/kaellybot/kaelly-webhooks/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	db, err := databases.New()
	if err != nil {
		return nil, err
	}

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		[]amqp.Binding{dispatchers.GetBinding()})
	if err != nil {
		return nil, err
	}

	// repositories
	webhooksRepo := webhooksRepo.New(db)
	feedsRepo := feedsRepo.New(db)

	// services
	webhooksService := webhooks.New(webhooksRepo)
	feedsService, err := feeds.New(feedsRepo)
	if err != nil {
		return nil, err
	}

	discordService, err := discord.New(viper.GetString(constants.Token))
	if err != nil {
		return nil, err
	}

	dispatchersService := dispatchers.New(broker, webhooksService, feedsService, discordService)

	return &Impl{
		broker:             broker,
		dispatchersService: dispatchersService,
		discordService:     discordService,
	}, nil
}

func (app *Impl) Run() error {
	return app.dispatchersService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.discordService.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
