package player

import (
	"context"
	"management-be/internal/model"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetPlayerStatistics(t *testing.T) {
	type args struct {
		id        int
		expResult model.PlayerStatistic
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			id: 1, // ID of the player in the database
			expResult: model.PlayerStatistic{
				PlayerID:           1,
				TotalMatches:       5,
				TotalMinutesPlayed: 450,
				TotalGoals:         3,
				TotalAssists:       2,
				TotalYellowCards:   1,
				TotalRedCards:      0,
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
					require.Equal(t, tc.expResult.PlayerID, statistics.PlayerID)
					require.Equal(t, tc.expResult.TotalMatches, statistics.TotalMatches)
					require.Equal(t, tc.expResult.TotalMinutesPlayed, statistics.TotalMinutesPlayed)
					require.Equal(t, tc.expResult.TotalGoals, statistics.TotalGoals)
					require.Equal(t, tc.expResult.TotalAssists, statistics.TotalAssists)
					require.Equal(t, tc.expResult.TotalYellowCards, statistics.TotalYellowCards)
					require.Equal(t, tc.expResult.TotalRedCards, statistics.TotalRedCards)
				}
			})
		})
	}
}
