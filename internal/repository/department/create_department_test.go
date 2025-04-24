package department

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestCreateDepartment(t *testing.T) {
	type args struct {
		name        string
		description string
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			name:        "New Department",
			description: "This is a new department",
		},
		"success - empty description": {
			name:        "Department with Empty Description",
			description: "",
		},
		"err - duplicate name": {
			name:        "Test Department", // Same as in insert_department.sql
			description: "Different description",
			expErr:      ErrDatabase,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				repo := NewRepository(tx.Client())
				department, err := repo.CreateDepartment(context.Background(), tc.name, tc.description)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotZero(t, department.ID)
					require.Equal(t, tc.name, department.Name)
					require.Equal(t, tc.description, department.Description)
					require.NotZero(t, department.CreatedAt)
					require.NotZero(t, department.UpdatedAt)
					// For model.Department, DeletedAt is a pointer that should be nil or point to a zero time
					if department.DeletedAt != nil {
						require.True(t, department.DeletedAt.IsZero())
					}

					// Verify the department was actually created in the database
					dbDepartment, err := tx.Client().Department.Get(context.Background(), department.ID)
					require.NoError(t, err)
					require.Equal(t, tc.name, dbDepartment.Name)
					require.Equal(t, tc.description, dbDepartment.Description)
					require.NotZero(t, dbDepartment.CreatedAt)
					require.NotZero(t, dbDepartment.UpdatedAt)
					// For ent.Department, DeletedAt is a time.Time value that should be zero
					require.True(t, dbDepartment.DeletedAt.IsZero())
				}
			})
		})
	}
}
