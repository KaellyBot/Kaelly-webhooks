package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
)

func MapAlmanax(almanax *amqp.NewsRSSMessage) *discordgo.WebhookParams {
	return &discordgo.WebhookParams{
		// TODO
	}
}
