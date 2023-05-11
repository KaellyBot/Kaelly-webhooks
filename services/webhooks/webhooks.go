package webhooks

import "github.com/kaellybot/kaelly-webhooks/repositories/webhooks"

func New(webhooksRepo webhooks.Repository) *Impl {
	return &Impl{webhooksRepo: webhooksRepo}
}
