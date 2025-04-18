package testent

import (
	"context"
	sqlconn "database/sql"
	"errors"
	"management-be/internal/pkg/config"
	"management-be/internal/repository/ent"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var appEntClient *ent.Client

// WithEntTx provides a callback with an `*ent.Tx` for running ent related tests
// where the `*ent.Tx` is actually powered by a database transaction
// and will be rolled back (so no data is actually written into the database)
func WithEntTx(t *testing.T, callback func(tx *ent.Tx)) {
	if appEntClient == nil {
		var err error
		// Replace with your actual database connection string
		drv, err := sql.Open(dialect.Postgres, os.Getenv("PG_URL"))

		// Create an ent.Client with the driver.
		appEntClient = ent.NewClient(ent.Driver(drv))

		// Set connection pool settings if needed
		db := drv.DB()
		db.SetMaxOpenConns(config.AppConfig.PgPoolMaxOpenConns)
		db.SetMaxIdleConns(config.AppConfig.PgPoolMaxIdleConns)
		db.SetConnMaxLifetime(15 * time.Minute)

		// Optional: Run migrations
		require.NoError(t, err)
	}

	ctx := context.Background()
	tx, err := appEntClient.Tx(ctx)
	require.NoError(t, err)

	defer func() {
		// Rollback transaction after callback to ensure no data is written
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sqlconn.ErrTxDone) {
			t.Errorf("Failed to rollback transaction: %v", err)
		}
	}()

	callback(tx)
}
