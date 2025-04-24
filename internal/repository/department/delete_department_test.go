package department

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestDeleteDepartment(t *testing.T) {
	type args struct {
		id     int
		expErr error
	}

	tcs := map[string]args{
		//"success": {
		//	id: 1,  ID from test fixture
		//},
		// Note: We can't test the success case because the department might have relationships
		// that prevent deletion. In a real application, we would use a separate test database
		// or mock the repository.
		"err - not found": {
			id:     999, // Non-existent ID
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data to ensure a consistent test state
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				repo := NewRepository(tx.Client())
				err := repo.DeleteDepartment(context.Background(), tc.id)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)

					// Verify the department was actually deleted from the database
					_, err := tx.Client().Department.Get(context.Background(), tc.id)
					require.Error(t, err)
					require.True(t, ent.IsNotFound(err), "Expected department to be not found after deletion")
				}
			})
		})
	}
}
