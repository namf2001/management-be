package match

import (
	"context"
	"testing"
	"time"

	"management-be/internal/model"
	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

// TestUpdateMatch tests the UpdateMatch function of the Match repository.
func TestUpdateMatch(t *testing.T) {
	type args struct {
		matchID   int
		input     UpdateMatchInput
		expResult model.Match
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			matchID: 1,
			input: UpdateMatchInput{
				OpponentTeamID: 2,
				MatchDate:      time.Date(2023, 10, 1, 15, 0, 0, 0, time.UTC),
				Venue:          "Updated Venue",
				IsHomeGame:     false,
				OurScore:       0,
				OpponentScore:  0,
				Status:         "completed",
				Notes:          "Updated Notes",
			},
			expResult: model.Match{
				ID:             1,
				OpponentTeamID: 2,
				MatchDate:      time.Date(2023, 10, 1, 15, 0, 0, 0, time.UTC),
				Venue:          "Updated Venue",
				IsHomeGame:     false,
				OurScore:       0,
				OpponentScore:  0,
				Status:         "completed",
				Notes:          "Updated Notes",
			},
			expErr: nil,
		},
		"err - match not found": {
			matchID: 999, // Non-existent match ID
			input: UpdateMatchInput{
				OpponentTeamID: 2,
				MatchDate:      time.Date(2023, 10, 1, 15, 0, 0, 0, time.UTC),
				Venue:          "Updated Venue",
				IsHomeGame:     false,
				OurScore:       0,
				OpponentScore:  0,
				Status:         "completed",
				Notes:          "Updated Notes",
			},
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Given
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")

				repo := NewRepository(tx.Client())
				match, err := repo.UpdateMatch(context.Background(), tc.matchID, tc.input)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expResult.ID, match.ID)
					require.Equal(t, tc.expResult.OpponentTeamID, match.OpponentTeamID)
					require.Equal(t, tc.expResult.MatchDate, match.MatchDate)
					require.Equal(t, tc.expResult.Venue, match.Venue)
					require.Equal(t, tc.expResult.IsHomeGame, match.IsHomeGame)
					require.Equal(t, tc.expResult.OurScore, match.OurScore)
					require.Equal(t, tc.expResult.OpponentScore, match.OpponentScore)
					require.Equal(t, tc.expResult.Status, match.Status)
					require.Equal(t, tc.expResult.Notes, match.Notes)

					require.NotZero(t, match.CreatedAt)
					require.NotZero(t, match.UpdatedAt)
				}
			})
		})
	}
}
