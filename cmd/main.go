package main

import (
	"context"
	"log"
	"os"
	"runtime"

	"github.com/thumperq/golib/config"
	"github.com/thumperq/wms/mailbox-cleanup-job/internal/application"
	"github.com/thumperq/wms/mailbox-cleanup-job/internal/infrastructure/db"
)

func main() {
	ctx := context.Background()
	configManager := config.NewConfigManager()
	dbFactory := db.NewDbFactory(configManager)
	appFactory := application.NewApplicationFactory(dbFactory)
	err := appFactory.MailboxApp.CleanupMailboxes(ctx)
	if err != nil {
		log.Printf("Error invoking the cleanup mailboxes: %v", err)
		os.Exit(1)
	}
	runtime.Goexit()
}
