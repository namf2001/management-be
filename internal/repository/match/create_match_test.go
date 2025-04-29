package match

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestCreateMatch(t *testing.T) {
	type args struct {
		opponentTeamID int
		matchDate      time.Time
		venue          string
		isHomeGame     bool
		notes          string
		expErr         error
	}

	tcs := map[string]args{
		"success": {
			opponentTeamID: 1,
			matchDate:      time.Now().UTC(), // Use UTC explicitly to match the database storage format
			venue:          "Test Venue",
			isHomeGame:     true,
			notes:          "Test Notes",
		},
		"err - invalid opponent team ID": {
			opponentTeamID: 0, // Invalid ID
			matchDate:      time.Now().UTC(), // Use UTC explicitly to match the database storage format
			venue:          "Test Venue",
			isHomeGame:     true,
			notes:          "Test Notes",
			expErr:         ErrDatabase,
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
				match, err := repo.CreateMatch(context.Background(), tc.opponentTeamID, tc.matchDate, tc.venue, tc.isHomeGame, tc.notes)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotZero(t, match.ID)
					require.Equal(t, tc.opponentTeamID, match.OpponentTeamID)
										require.Equal(t, tc.matchDate, match.MatchDate)
					require.Equal(t, tc.venue, match.Venue)
					require.Equal(t, tc.isHomeGame, match.IsHomeGame)
					require.Equal(t, tc.notes, match.Notes)

					// Verify the match was actually created in the database
					dbMatch, err := tx.Client().Match.Get(context.Background(), match.ID)
					require.NoError(t, err)
					require.Equal(t, tc.opponentTeamID, dbMatch.OpponentTeamID)
								require.Equal(t, tc.matchDate, dbMatch.MatchDate)
					require.Equal(t, tc.venue, dbMatch.Venue)
					require.Equal(t, tc.isHomeGame, dbMatch.IsHomeGame)
					require.Equal(t, tc.notes, dbMatch.Notes)
				}
			})
		})
	}
}
