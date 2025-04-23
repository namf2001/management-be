package department

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetDepartmentByID(t *testing.T) {
	type args struct {
		id     int
		expErr error
	}

	tcs := map[string]args{
		"success": {
			id: 1, // ID of the department in the database
		},
		"err - not found": {
			id:     999, // Non-existent ID
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				repo := NewRepository(tx.Client())
				department, err := repo.GetDepartmentByID(context.Background(), tc.id)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.id, department.ID)
					require.Equal(t, "Engineering", department.Name) // Use the actual name from the database
					require.NotEmpty(t, department.Description)
					require.NotZero(t, department.CreatedAt)
					require.NotZero(t, department.UpdatedAt)
				}
			})
		})
	}
}
