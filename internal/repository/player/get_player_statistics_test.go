package player

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetPlayerStatistics(t *testing.T) {
	type args struct {
		id     int
		expErr error
	}

	tcs := map[string]args{
		"success": {
			id: 1, // ID of the player in the database
		},
		"err - not found": {
			id:     999, // Non-existent ID
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data to ensure department, player, and statistics exist
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player_statistics.sql")

				repo := NewRepository(tx.Client())
				statistics, err := repo.GetPlayerStatistics(context.Background(), tc.id)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.id, statistics.PlayerID)
					require.Equal(t, int32(10), statistics.TotalMatches)
					require.Equal(t, int32(900), statistics.TotalMinutesPlayed)
					require.Equal(t, int32(5), statistics.TotalGoals)
					require.Equal(t, int32(3), statistics.TotalAssists)
					require.Equal(t, int32(2), statistics.TotalYellowCards)
					require.Equal(t, int32(0), statistics.TotalRedCards)
					require.NotZero(t, statistics.CreatedAt)
					require.NotZero(t, statistics.UpdatedAt)
				}
			})
		})
	}
}
