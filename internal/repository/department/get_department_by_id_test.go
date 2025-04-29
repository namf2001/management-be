package department

import (
	"context"
	"testing"

	"management-be/internal/model"
	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetDepartmentByID(t *testing.T) {
	type args struct {
		id        int
		expResult model.Department
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			id: 1, // ID of the department in the database
			expResult: model.Department{
				ID:          1,
				Name:        "Test Department",
				Description: "This is a test department",
			},
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
					require.Equal(t, tc.expResult.ID, department.ID)
					require.Equal(t, tc.expResult.Name, department.Name)
					require.Equal(t, tc.expResult.Description, department.Description)
					require.NotZero(t, department.CreatedAt)
					require.NotZero(t, department.UpdatedAt)
				}
			})
		})
	}
}
