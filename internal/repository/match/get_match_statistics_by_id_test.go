package match

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetMatchStatistics(t *testing.T) {
	type args struct {
		matchID int
		expErr  error
	}

	tcs := map[string]args{
		"success - match with players": {
			matchID: 1,
		},
		"success - match without players": {
			matchID: 2, // Match exists but has no players associated
		},
		"err - match not found": {
			matchID: 999, // Non-existent match ID
			expErr:  ErrDatabase,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match_players.sql")

				repo := NewRepository(tx.Client())
				stats, err := repo.GetMatchStatistics(context.Background(), tc.matchID)

				// then
				if tc.expErr != nil {
					require.Error(t, err)
				} else {
					require.NoError(t, err)

					if tc.matchID == 1 {
						// For match 1, we expect 3 players with specific stats
						require.Equal(t, int32(3), stats.MatchSummary.TotalPlayers)
						require.Equal(t, int32(90+85+90), stats.MatchSummary.TotalMinutesPlayed)
						require.Equal(t, int32(2), stats.MatchSummary.TotalGoals)
						require.Equal(t, int32(3), stats.MatchSummary.TotalAssists)
						require.Equal(t, int32(1), stats.MatchSummary.TotalYellowCards)
						require.Equal(t, int32(1), stats.MatchSummary.TotalRedCards)

						require.Len(t, stats.PlayerPerformance, 3)

						// Verify at least one player performance entry (Player One with 2 goals)
						foundPlayer := false
						for _, player := range stats.PlayerPerformance {
							if player.PlayerID == 1 {
								foundPlayer = true
								require.Equal(t, "Player One", player.PlayerName)
								require.Equal(t, "Forward", player.Position)
								require.Equal(t, int32(90), player.MinutesPlayed)
								require.Equal(t, int32(2), player.GoalsScored)
								require.Equal(t, int32(1), player.Assists)
								require.Equal(t, int32(0), player.YellowCards)
								require.False(t, player.RedCard)
								break
							}
						}
						require.True(t, foundPlayer, "Player One should be found in the statistics")
					} else if tc.matchID == 2 {
						// For match 2, we expect no players
						require.Equal(t, int32(0), stats.MatchSummary.TotalPlayers)
						require.Equal(t, int32(0), stats.MatchSummary.TotalMinutesPlayed)
						require.Equal(t, int32(0), stats.MatchSummary.TotalGoals)
						require.Equal(t, int32(0), stats.MatchSummary.TotalAssists)
						require.Equal(t, int32(0), stats.MatchSummary.TotalYellowCards)
						require.Equal(t, int32(0), stats.MatchSummary.TotalRedCards)
						require.Empty(t, stats.PlayerPerformance)
					}
				}
			})
		})
	}
}
