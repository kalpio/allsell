package migrations

import (
	"context"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	_ "modernc.org/sqlite"
)

type MigrationConfig struct {
	Driver           string
	ConnectionString string
}

func Migrate(ctx context.Context, config MigrationConfig) error {
	db, err := goose.OpenDBWithDriver(config.Driver, config.ConnectionString)
	if err != nil {
		return fmt.Errorf("migration: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(fmt.Errorf("migration: %w", err))
		}
	}()

	if err = goose.RunContext(ctx, "up", db, ".", []string{}...); err != nil {
		return fmt.Errorf("migration: %w", err)
	}

	return nil
}
