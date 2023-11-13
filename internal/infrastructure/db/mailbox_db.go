package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/thumperq/golib/database"
)

type MailboxDB interface {
	DeleteMailboxes(ctx context.Context) error
}
type MailboxDb struct {
	pgDb database.PgDB
}

func NewMailboxDb(pgDb database.PgDB) MailboxDB {
	return MailboxDb{
		pgDb: pgDb,
	}
}

func (db MailboxDb) DeleteMailboxes(ctx context.Context) error {
	err := db.pgDb.WithTransaction(ctx, func(tx pgx.Tx) error {
		sql, args, err := sb.Delete("mailboxes").ToSql()
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
