package testent

import (
	"context"
	sqlconn "database/sql"
	"errors"
	"fmt"
	"management-be/internal/pkg/config"
	"management-be/internal/repository/ent"
	"os"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var appEntClient *ent.Client

// setTestEnv sets up the environment variables needed for testing
func setTestEnv() {
	// Manually set environment variables for testing
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "management-football")
	os.Setenv("DB_PASSWORD", "management-football")
	os.Setenv("DB_NAME", "management-football")
	os.Setenv("DB_SCHEMA", "public")
}

// buildConnString constructs the database connection string
func buildConnString() string {
	// Replace with your actual database connection string
	pgURL := os.Getenv("PG_URL")
	if pgURL != "" {
		// Ensure sslmode=disable is set
		if !strings.Contains(pgURL, "sslmode=") {
			if strings.Contains(pgURL, "?") {
				pgURL += "&sslmode=disable"
			} else {
				pgURL += "?sslmode=disable"
			}
		}
		return pgURL
	}

	// Construct the connection string from individual environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		user, password, host, port, dbname, schema)
}

// WithEntTx provides a callback with an `*ent.Tx` for running ent related tests
// where the `*ent.Tx` is actually powered by a database transaction
// and will be rolled back (so no data is actually written into the database)
func WithEntTx(t *testing.T, callback func(tx *ent.Tx)) {
	// Skip tests in CI/CD environment
	if os.Getenv("CI") != "" {
		t.Skip("Skipping database tests in CI environment")
		return
	}

	if appEntClient == nil {
		setTestEnv()
		drv, err := sql.Open(dialect.Postgres, buildConnString())
		require.NoError(t, err)

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
