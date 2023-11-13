package application

import (
	"github.com/thumperq/wms/mailbox-cleanup-job/internal/infrastructure/db"
)

type AppFactory struct {
	MailboxApp MailboxApp
}

func NewApplicationFactory(DbFactory db.DbFactory) AppFactory {
	return AppFactory{
		MailboxApp: NewMailboxApp(DbFactory),
	}
}
