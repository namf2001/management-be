package user

import (
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	testent.WithEntTx(t, func(tx *ent.Tx) {
		// Create a new repository with the Ent client
		repo := NewRepository(tx.Client())

		// Verify that the repository is not nil
		require.NotNil(t, repo)

		// Verify that the repository implements the Repository interface
		_, ok := repo.(Repository)
		require.True(t, ok, "Repository should implement the Repository interface")
	})
}
