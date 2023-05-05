package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	broker amqp.MessageBroker
}
