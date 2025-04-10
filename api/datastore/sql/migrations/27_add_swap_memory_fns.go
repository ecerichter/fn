package migrations

import (
	"context"

	"github.com/fnproject/fn/api/datastore/sql/migratex"
	"github.com/jmoiron/sqlx"
)

func up27(ctx context.Context, tx *sqlx.Tx) error {

	// Firstly add a new column shape
	_, err := tx.ExecContext(ctx, "ALTER TABLE fns ADD swap_memory INTEGER NOT NULL DEFAULT 0;")
	if err != nil {
		return err
	}

	return nil
}

func down27(ctx context.Context, tx *sqlx.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE fns DROP COLUMN swap_memory;")
	return err
}

func init() {
	Migrations = append(Migrations, &migratex.MigFields{
		VersionFunc: vfunc(27),
		UpFunc:      up27,
		DownFunc:    down27,
	})
}
