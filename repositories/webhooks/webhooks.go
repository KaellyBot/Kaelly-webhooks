package webhooks

import "github.com/kaellybot/kaelly-webhooks/utils/databases"

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}
