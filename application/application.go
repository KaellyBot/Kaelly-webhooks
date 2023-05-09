package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	almanaxRepo "github.com/kaellybot/kaelly-webhooks/repositories/almanax"
	feedsRepo "github.com/kaellybot/kaelly-webhooks/repositories/feeds"
	twitterRepo "github.com/kaellybot/kaelly-webhooks/repositories/twitter"
	"github.com/kaellybot/kaelly-webhooks/services/almanax"
	"github.com/kaellybot/kaelly-webhooks/services/discord"
	"github.com/kaellybot/kaelly-webhooks/services/feeds"
	"github.com/kaellybot/kaelly-webhooks/services/twitter"
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
		[]amqp.Binding{webhooks.GetBinding()})
	if err != nil {
		return nil, err
	}

	// repositories
	almanaxRepo := almanaxRepo.New(db)
	feedsRepo := feedsRepo.New(db)
	twitterRepo := twitterRepo.New(db)

	// services
	almanaxService := almanax.New(almanaxRepo)
	feedsService := feeds.New(feedsRepo)
	twitterService := twitter.New(twitterRepo)
	discordService, err := discord.New(viper.GetString(constants.Token))
	if err != nil {
		return nil, err
	}

	webhooksService := webhooks.New(broker, almanaxService, feedsService,
		twitterService, discordService)

	return &Impl{
		broker:          broker,
		webhooksService: webhooksService,
		discordService:  discordService,
	}, nil
}

func (app *Impl) Run() error {
	return app.webhooksService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.discordService.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
