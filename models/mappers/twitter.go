package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
)

func MapTweet(tweet *amqp.NewsTwitterMessage) *discordgo.WebhookParams {
	return &discordgo.WebhookParams{
		Content: tweet.Url,
	}
}
