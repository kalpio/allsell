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
	query := `CREATE TABLE users (
id TEXT PRIMARY KEY,
name TEXT NOT NULL UNIQUE,
email TEXT NOT NULL UNIQUE,
password TEXT,
last_password_change TEXT,
last_login_at TEXT);`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func down0001(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE users;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
