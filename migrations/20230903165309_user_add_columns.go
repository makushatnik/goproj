package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserAddColumns, downUserAddColumns)
}

func upUserAddColumns(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users ADD COLUMN last_name VARCHAR(100) NOT NULL DEFAULT '';
		ALTER TABLE users ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();
		ALTER TABLE users ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT NOW();
	`)
	if err != nil {
		return err
	}
	return nil
}

func downUserAddColumns(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users DROP COLUMN last_name;
		ALTER TABLE users DROP COLUMN created_at;
		ALTER TABLE users DROP COLUMN updated_at;
	`)
	if err != nil {
		return err
	}
	return nil
}
