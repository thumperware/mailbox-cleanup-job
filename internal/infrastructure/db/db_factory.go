package db

import (
	"github.com/Masterminds/squirrel"
	"github.com/thumperq/golib/config"
	"github.com/thumperq/golib/database"
)

var sb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type DbFactory struct {
	MailboxDb MailboxDB
}

func NewDbFactory(cfg config.CfgManager) DbFactory {
	pgDb := database.NewPostgresConnection(cfg)
	return DbFactory{
		MailboxDb: NewMailboxDb(pgDb),
	}
}
