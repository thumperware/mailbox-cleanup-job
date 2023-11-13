package application

import (
	"context"

	"github.com/thumperq/wms/mailbox-cleanup-job/internal/infrastructure/db"
)

type MailboxApp struct {
	dbFactory db.DbFactory
}

func NewMailboxApp(DbFactory db.DbFactory) MailboxApp {
	app := MailboxApp{
		dbFactory: DbFactory,
	}
	return app
}

func (app MailboxApp) CleanupMailboxes(ctx context.Context) error {
	return app.dbFactory.MailboxDb.DeleteMailboxes(ctx)
}
