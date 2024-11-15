package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up0003, down0003)
}

func up0003(ctx context.Context, tx *sql.Tx) error {
	query := `create table auctions (
id TEXT NOT NULL PRIMARY KEY,
title TEXT NOT NULL,
expire_at TEXT NOT NULL,
category TEXT NOT NULL,
price TEXT NOT NULL
);`
	_, err := tx.ExecContext(ctx, query)

	return err
}

func down0003(ctx context.Context, tx *sql.Tx) error {
	query := `drop table auctions;`

	_, err := tx.ExecContext(ctx, query)
	return err
}
