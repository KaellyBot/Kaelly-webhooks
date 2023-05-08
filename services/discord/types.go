package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Service interface {
	PublishWebhook(webhookID, webhookToken string, data *discordgo.WebhookParams) error
	Shutdown()
}

type Impl struct {
	session *discordgo.Session
}
