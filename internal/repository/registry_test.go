package repository

import (
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestNewRegistry(t *testing.T) {
	testent.WithEntTx(t, func(tx *ent.Tx) {
		// Create a new registry with the Ent client
		registry := NewRegistry(tx.Client())

		// Verify that the registry is not nil
		require.NotNil(t, registry)

		// Verify that the registry implements the Registry interface
		_, ok := registry.(Registry)
		require.True(t, ok, "Registry should implement the Registry interface")

		// Verify that the User() method returns a non-nil user repository
		userRepo := registry.User()
		require.NotNil(t, userRepo)
	})
}
