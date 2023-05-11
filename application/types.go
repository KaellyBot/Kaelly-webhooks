package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/services/discord"
	"github.com/kaellybot/kaelly-webhooks/services/dispatchers"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	broker             amqp.MessageBroker
	dispatchersService dispatchers.Service
	discordService     discord.Service
}
