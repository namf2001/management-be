package department

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetAllDepartments(t *testing.T) {
	t.Run("success - get all departments", func(t *testing.T) {
		testent.WithEntTx(t, func(tx *ent.Tx) {
			// No need to load test data, as there are already departments in the database

			repo := NewRepository(tx.Client())
			departments, err := repo.GetAllDepartments(context.Background())

			require.NoError(t, err)
			require.NotEmpty(t, departments)

			// Verify that departments have the expected structure
			for _, dept := range departments {
				require.NotZero(t, dept.ID)
				require.NotEmpty(t, dept.Name)
				require.NotZero(t, dept.CreatedAt)
				require.NotZero(t, dept.UpdatedAt)
			}
		})
	})

	// Note: We can't test the "no departments" case because there are already departments in the test database
	// and we can't delete them in a transaction that will be rolled back.
	// In a real application, we would use a separate test database or mock the repository.
}
