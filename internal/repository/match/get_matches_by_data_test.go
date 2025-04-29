package match

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestListMatches(t *testing.T) {
	type args struct {
		status         string
		startDate      time.Time
		endDate        time.Time
		opponentTeamID int
		expectedCount  int
	}

	defaultStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	defaultEndDate := time.Date(2025, 5, 31, 23, 59, 59, 0, time.UTC)

	tcs := map[string]args{
		"list all matches": {
			status:         "",
			startDate:      time.Time{},
			endDate:        time.Time{},
			opponentTeamID: 0,
			expectedCount:  2,
		},
		"filter by status": {
			status:         "scheduled",
			startDate:      time.Time{},
			endDate:        time.Time{},
			opponentTeamID: 0,
			expectedCount:  2,
		},
		"filter by date range": {
			status:         "",
			startDate:      defaultStartDate,
			endDate:        defaultEndDate,
			opponentTeamID: 0,
			expectedCount:  2,
		},
		"filter by opponent team": {
			status:         "",
			startDate:      time.Time{},
			endDate:        time.Time{},
			opponentTeamID: 1,
			expectedCount:  1,
		},
		"filter combinations": {
			status:         "scheduled",
			startDate:      defaultStartDate,
			endDate:        defaultEndDate,
			opponentTeamID: 1,
			expectedCount:  1,
		},
		"no matches found": {
			status:         "completed", // No completed matches in test data
			startDate:      time.Time{},
			endDate:        time.Time{},
			opponentTeamID: 0,
			expectedCount:  0,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Clean the tables first to avoid conflicts
				_, err := tx.ExecContext(context.Background(), "DELETE FROM matches")
				require.NoError(t, err)
				_, err = tx.ExecContext(context.Background(), "DELETE FROM teams")
				require.NoError(t, err)
				
				// Load team data first, then match data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")

				repo := NewRepository(tx.Client())
				matches, err := repo.ListMatches(context.Background(), tc.status, tc.startDate, tc.endDate, tc.opponentTeamID)

				// then
				require.NoError(t, err)
				require.Len(t, matches, tc.expectedCount)

				// If expecting matches, verify they have the expected properties
				if tc.expectedCount > 0 {
					for _, match := range matches {
						// Verify match properties based on filters
						if tc.status != "" {
							require.Equal(t, tc.status, match.Status)
						}
						if !tc.startDate.IsZero() {
							require.GreaterOrEqual(t, match.MatchDate.Unix(), tc.startDate.Unix())
						}
						if !tc.endDate.IsZero() {
							require.LessOrEqual(t, match.MatchDate.Unix(), tc.endDate.Unix())
						}
						if tc.opponentTeamID > 0 {
							require.Equal(t, tc.opponentTeamID, match.OpponentTeamID)
						}
					}
				}
			})
		})
	}
}