package application

import (
	"fmt"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
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

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress), nil)
	if err != nil {
		return nil, err
	}

	// repositories
	fmt.Printf("%v", db)
	// TODO

	// services
	// TODO

	return &Impl{broker: broker}, nil
}

func (app *Impl) Run() error {
	// TODO
	return nil
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
