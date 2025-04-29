package match

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestUpdateMatch(t *testing.T) {
	type args struct {
		matchID        int
		opponentTeamID int
		matchDate      time.Time
		venue          string
		isHomeGame     bool
		ourScore       int32
		opponentScore  int32
		status         string
		notes          string
		expErr         error
	}

	tcs := map[string]args{
		"success": {
			matchID:        1,
			opponentTeamID: 2,
			matchDate:      time.Now().Add(24 * time.Hour),
			venue:          "Updated Venue",
			isHomeGame:     false,
			ourScore:       2,
			opponentScore:  1,
			status:         "completed",
			notes:          "Updated Notes",
		},
		"err - match not found": {
			matchID:        999, // Non-existent match ID
			opponentTeamID: 2,
			matchDate:      time.Now().Add(24 * time.Hour),
			venue:          "Updated Venue",
			isHomeGame:     false,
			ourScore:       0,
			opponentScore:  0,
			status:         "completed",
			notes:          "Updated Notes",
			expErr:         ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Given
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")

				repo := NewRepository(tx.Client())
				match, err := repo.UpdateMatch(context.Background(), tc.matchID, tc.opponentTeamID, tc.matchDate, tc.venue, tc.isHomeGame, tc.ourScore, tc.opponentScore, tc.status, tc.notes)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotNil(t, match)
					require.Equal(t, tc.matchID, match.ID)
					require.Equal(t, tc.venue, match.Venue)
					require.Equal(t, tc.status, match.Status)
					require.Equal(t, tc.ourScore, match.OurScore)
					require.Equal(t, tc.opponentScore, match.OpponentScore)
				}
			})
		})
	}
}
