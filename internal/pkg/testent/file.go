package testent

import (
	"context"
	"management-be/internal/repository/ent"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// LoadTestSQLFile loads a SQL file from the filesystem and executes it in the given ent.Tx.
// It mainly serves for testing purpose
func LoadTestSQLFile(t *testing.T, entTx *ent.Tx, filename string) {
	cwd, err := os.Getwd()
	require.NoError(t, err)
	fullpath, err := filepath.Abs(filepath.Join(cwd, filename))
	require.NoError(t, err)
	body, err := os.ReadFile(fullpath)
	require.NoError(t, err)

	_, err = entTx.ExecContext(context.Background(), string(body))
	require.NoError(t, err)
}
