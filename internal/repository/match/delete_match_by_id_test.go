package match

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestDeleteMatchByID(t *testing.T) {
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
				// Clean the tables first to avoid conflicts
				_, err := tx.ExecContext(context.Background(), "DELETE FROM matches")
				require.NoError(t, err)
				_, err = tx.ExecContext(context.Background(), "DELETE FROM teams")
				require.NoError(t, err)
				
				// Load team data first, then match data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_match.sql")

				repo := NewRepository(tx.Client())
				err = repo.DeleteMatchByID(context.Background(), tc.matchID)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					
					// Verify the match was actually deleted
					_, err := tx.Client().Match.Get(context.Background(), tc.matchID)
					require.Error(t, err, "Match should not exist after deletion")
				}
			})
		})
	}
}