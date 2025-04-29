package player

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/matchplayer"
	"management-be/internal/repository/ent/playerstatistic"

	"github.com/stretchr/testify/require"
)

func TestDeletePlayer(t *testing.T) {
	type args struct {
		id                      int
		hasMatchPlayerData      bool
		hasPlayerStatisticsData bool
		expErr                  error
	}

	tcs := map[string]args{
		"success": {
			id: 1, // ID of the player in the database
		},
		"success - with match player data": {
			id:                 1,
			hasMatchPlayerData: true,
		},
		"success - with player statistics data": {
			id:                      1,
			hasPlayerStatisticsData: true,
		},
		"success - with both related data": {
			id:                      1,
			hasMatchPlayerData:      true,
			hasPlayerStatisticsData: true,
		},
		"err - not found": {
			id:     999, // Non-existent ID
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// First, insert department data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				// Then, insert player data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

				// Insert match player data if needed
				if tc.hasMatchPlayerData {
					testent.LoadTestSQLFile(t, tx, "testdata/insert_match_player.sql")

					// Verify match player data exists before deletion
					count, err := tx.Client().MatchPlayer.Query().
						Where(matchplayer.PlayerID(tc.id)).
						Count(context.Background())
					require.NoError(t, err)
					require.Greater(t, count, 0, "Expected match player data to exist before deletion")
				}

				// Insert player statistics data if needed
				if tc.hasPlayerStatisticsData {
					testent.LoadTestSQLFile(t, tx, "testdata/insert_player_statistics.sql")

					// Verify player statistics data exists before deletion
					count, err := tx.Client().PlayerStatistic.Query().
						Where(playerstatistic.PlayerID(tc.id)).
						Count(context.Background())
					require.NoError(t, err)
					require.Greater(t, count, 0, "Expected player statistics data to exist before deletion")
				}

				repo := NewRepository(tx.Client())
				err := repo.DeletePlayer(context.Background(), tc.id)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)

					// Verify the player was actually deleted from the database
					_, err := tx.Client().Player.Get(context.Background(), tc.id)
					require.Error(t, err)
					require.True(t, ent.IsNotFound(err), "Expected player to be not found after deletion")

					// For cases with match player data, verify those records were deleted too
					if tc.hasMatchPlayerData {
						count, err := tx.Client().MatchPlayer.Query().
							Where(matchplayer.PlayerID(tc.id)).
							Count(context.Background())
						require.NoError(t, err)
						require.Equal(t, 0, count, "Expected all match player records to be deleted")
					}

					// For cases with player statistics data, verify those records were deleted too
					if tc.hasPlayerStatisticsData {
						count, err := tx.Client().PlayerStatistic.Query().
							Where(playerstatistic.PlayerID(tc.id)).
							Count(context.Background())
						require.NoError(t, err)
						require.Equal(t, 0, count, "Expected all player statistics records to be deleted")
					}
				}
			})
		})
	}
}
