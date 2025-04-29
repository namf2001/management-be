package match

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetMatchByID(t *testing.T) {
	type args struct {
		matchID int
		expErr  error
	}

	tcs := map[string]args{
		"success": {
			matchID: 1,
		},
		"err - match not found": {
			matchID: 999, // Non-existent match ID
			expErr:  ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Given
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")

				repo := NewRepository(tx.Client())
				match, err := repo.GetMatchByID(context.Background(), tc.matchID)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotNil(t, match)
					require.Equal(t, tc.matchID, match.ID)
				}
			})
		})
	}
}
