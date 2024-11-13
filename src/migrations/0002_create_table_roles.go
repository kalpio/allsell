package migrations

import (
	"context"
	"database/sql"
	"github.com/kalpio/allsell/src/types/role"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up0002, down0002)
}

func up0002(ctx context.Context, tx *sql.Tx) error {
	query := `create table roles (
id TEXT PRIMARY KEY,
name TEXT NOT NULL UNIQUE
);`
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	query = `create table user_roles (
user_id TEXT NOT NULL,
role_id TEXT NOT NULL,
foreign key (user_id) references users(id),
foreign key (role_id) references roles(id),
unique (user_id, role_id) on conflict REPLACE
);`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	query = `insert into roles (id, name) values
(?, ?),
(?, ?);`

	_, err := tx.ExecContext(
		ctx, query,
		role.Administrator.ID, role.Administrator.Name,
		role.User.ID, role.User.Name)

	return err
}

func down0002(ctx context.Context, tx *sql.Tx) error {
	query := `drop table roles;`
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	query = `drop table user_roles;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
