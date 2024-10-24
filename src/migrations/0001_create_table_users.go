package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up0001, down0001)
}

func up0001(ctx context.Context, tx *sql.Tx) error {
	sql := `CREATE TABLE users (
id TEXT PRIMARY KEY,
name TEXT NOT NULL UNIQUE,
email TEXT NOT NULL UNIQUE,
password TEXT);`
	_, err := tx.ExecContext(ctx, sql)
	return err
}

func down0001(ctx context.Context, tx *sql.Tx) error {
	sql := `DROP TABLE users;`
	_, err := tx.ExecContext(ctx, sql)
	return err
}
