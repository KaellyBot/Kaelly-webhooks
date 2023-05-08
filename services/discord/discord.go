package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func New(token string) (*Impl, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Error().Err(err).Msgf("Connecting to Discord gateway failed")
		return nil, err
	}

	service := Impl{
		session: dg,
	}

	return &service, nil
}

func (service *Impl) PublishWebhook(webhookID, webhookToken string, data *discordgo.WebhookParams) error {
	_, err := service.session.WebhookExecute(webhookID, webhookToken, false, data)
	return err
}

func (service *Impl) Shutdown() {
	log.Info().Msgf("Closing Discord connections...")
	err := service.session.Close()
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot close session and shutdown correctly")
	}
}
