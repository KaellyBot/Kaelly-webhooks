package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type ApplicationInterface interface {
	Run() error
	Shutdown()
}

type Application struct {
	broker amqp.MessageBrokerInterface
}
