package department

import (
	"context"
	"errors"
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
			name:        "Test Department",
			description: "Different description",
			expErr:      errors.New("duplicate key value violates unique constraint \"unique_department_name\""),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {

				testent.LoadTestSQLFile(t, tx, "testdata/insert_department_without_alter.sql")

				repo := NewRepository(tx.Client())
				department, err := repo.CreateDepartment(context.Background(), tc.name, tc.description)

				if tc.expErr != nil {
					require.Error(t, err)
					require.Contains(t, err.Error(), tc.expErr.Error())
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
