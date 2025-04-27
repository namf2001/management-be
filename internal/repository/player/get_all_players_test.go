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
		page     int
		pageSize int
		expErr   error
	}

	tcs := map[string]args{
		"success - get all players": {
			page:     1,
			pageSize: 10,
		},
		"success - empty page": {
			page:     2,
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
				players, total, err := repo.GetAllPlayers(context.Background(), tc.page, tc.pageSize)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotNil(t, players)
					require.GreaterOrEqual(t, total, 0)

					if tc.page == 1 {
						require.Greater(t, len(players), 0)
					} else {
						require.Equal(t, 0, len(players))
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
