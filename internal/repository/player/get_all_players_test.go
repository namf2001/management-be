package player

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetAllPlayers(t *testing.T) {
	type args struct {
		page         int
		pageSize     int
		departmentID *int
		isActive     *bool
		position     string
		expErr       error
		expectEmpty  bool // New field to indicate if we expect empty results
	}

	tcs := map[string]args{
		"success - get all players": {
			page:     1,
			pageSize: 10,
		},
		"success - empty page": {
			page:        2,
			pageSize:    10,
			expectEmpty: true,
		},
		"success - filter by department": {
			page:         1,
			pageSize:     10,
			departmentID: ptrInt(1),
		},
		"success - filter by active status": {
			page:     1,
			pageSize: 10,
			isActive: ptrBool(true),
		},
		"success - filter by position": {
			page:     1,
			pageSize: 10,
			position: "Forward",
		},
		"success - invalid page number (defaults to page 1)": {
			page:     0,
			pageSize: 10,
		},
		"success - negative page number (defaults to page 1)": {
			page:     -1,
			pageSize: 10,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

				repo := NewRepository(tx.Client())
				players, total, err := repo.GetAllPlayers(
					context.Background(),
					tc.page,
					tc.pageSize,
					tc.departmentID,
					tc.isActive,
					tc.position,
				)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotNil(t, players)
					require.GreaterOrEqual(t, total, 0)

					if tc.expectEmpty {
						require.Equal(t, 0, len(players))
					} else {
						require.Greater(t, len(players), 0)
					}

					// Verify each player has required fields
					for _, player := range players {
						require.NotZero(t, player.ID)
						require.NotZero(t, player.DepartmentID)
						require.NotEmpty(t, player.FullName)
						require.NotEmpty(t, player.Position)
						require.NotNil(t, player.JerseyNumber)
						require.NotNil(t, player.DateOfBirth)
						require.NotNil(t, player.HeightCM)
						require.NotNil(t, player.WeightKG)
						require.NotEmpty(t, player.Phone)
						require.NotEmpty(t, player.Email)
						require.NotZero(t, player.CreatedAt)
						require.NotZero(t, player.UpdatedAt)
					}
				}
			})
		})
	}
}

// Helper functions to create pointers
func ptrInt(i int) *int {
	return &i
}

func ptrBool(b bool) *bool {
	return &b
}
