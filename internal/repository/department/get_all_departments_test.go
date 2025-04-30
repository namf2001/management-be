package department

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetAllDepartments(t *testing.T) {
	type args struct {
		expResult []ent.Department
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			expResult: []ent.Department{
				{
					ID:          1,
					Name:        "Test Department",
					Description: "This is a test department",
				},
			},
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				repo := NewRepository(tx.Client())
				departments, err := repo.GetAllDepartments(context.Background())

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, len(tc.expResult), len(departments))
					for i := range departments {
						require.Equal(t, tc.expResult[i].ID, departments[i].ID)
						require.Equal(t, tc.expResult[i].Name, departments[i].Name)
						require.Equal(t, tc.expResult[i].Description, departments[i].Description)
						require.NotZero(t, departments[i].CreatedAt)
						require.NotZero(t, departments[i].UpdatedAt)
					}
				}
			})
		})
	}
}
