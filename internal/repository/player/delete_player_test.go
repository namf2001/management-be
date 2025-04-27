package player

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestDeletePlayer(t *testing.T) {
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
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

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
				}
			})
		})
	}
}
