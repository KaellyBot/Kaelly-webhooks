package mappers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-webhooks/models/constants"
	"github.com/kaellybot/kaelly-webhooks/services/feeds"
)

func MapFeed(feed *amqp.NewsRSSMessage, feedService feeds.Service, locale amqp.Language) *discordgo.WebhookParams {
	feedLabel := feed.Type
	feedType, found := feedService.FindFeedTypeByID(feed.Type)
	if found {
		label, labelFound := feedType.GetLabels()[locale]
		if labelFound {
			feedLabel = label
		}
	}

	return &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: feed.Title,
				Author: &discordgo.MessageEmbedAuthor{
					Name: feed.AuthorName,
				},
				Color: constants.RSSColor,
				URL:   feed.Url,
				Image: &discordgo.MessageEmbedImage{
					URL: feed.IconUrl,
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: constants.RSSIconURL,
				},
				Timestamp: feed.Date.AsTime().Format(time.RFC3339),
				Footer: &discordgo.MessageEmbedFooter{
					Text: feedLabel,
				},
			},
		},
	}
}
