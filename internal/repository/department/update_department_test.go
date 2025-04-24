package department

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestUpdateDepartment(t *testing.T) {
	type args struct {
		id          int
		name        string
		description string
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			id:          1, // ID of the department inserted in insert_department.sql
			name:        "Updated Department",
			description: "This is an updated department",
		},
		"success - empty description": {
			id:          1, // ID of the department inserted in insert_department.sql
			name:        "Department with Empty Description",
			description: "",
		},
		"err - not found": {
			id:          999, // Non-existent ID
			name:        "Non-existent Department",
			description: "This department doesn't exist",
			expErr:      ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				// Record the time before update to verify updated_at is changed
				beforeUpdate := time.Now()
				time.Sleep(10 * time.Millisecond) // Ensure time difference

				repo := NewRepository(tx.Client())
				department, err := repo.UpdateDepartment(context.Background(), tc.id, tc.name, tc.description)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.id, department.ID)
					require.Equal(t, tc.name, department.Name)
					require.Equal(t, tc.description, department.Description)
					require.NotZero(t, department.CreatedAt)
					require.NotZero(t, department.UpdatedAt)
					require.True(t, department.UpdatedAt.After(beforeUpdate), "UpdatedAt should be after the time before update")

					// Verify the department was actually updated in the database
					dbDepartment, err := tx.Client().Department.Get(context.Background(), tc.id)
					require.NoError(t, err)
					require.Equal(t, tc.name, dbDepartment.Name)
					require.Equal(t, tc.description, dbDepartment.Description)
					require.NotZero(t, dbDepartment.CreatedAt)
					require.NotZero(t, dbDepartment.UpdatedAt)
					require.True(t, dbDepartment.UpdatedAt.After(beforeUpdate), "UpdatedAt in DB should be after the time before update")
				}
			})
		})
	}
}
