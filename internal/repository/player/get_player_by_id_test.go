package player

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetPlayerByID(t *testing.T) {
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
				// Load test data to ensure department and player exist
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

				repo := NewRepository(tx.Client())
				player, err := repo.GetPlayerByID(context.Background(), tc.id)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.id, player.ID)
					require.Equal(t, 1, player.DepartmentID)
					require.Equal(t, "Test Player", player.FullName)
					require.Equal(t, "Forward", player.Position)
					require.Equal(t, int32(10), *player.JerseyNumber)
					require.Equal(t, int32(180), *player.HeightCM)
					require.Equal(t, int32(75), *player.WeightKG)
					require.Equal(t, "1234567890", player.Phone)
					require.Equal(t, "test@example.com", player.Email)
					require.True(t, player.IsActive)
					require.NotZero(t, player.CreatedAt)
					require.NotZero(t, player.UpdatedAt)
				}
			})
		})
	}
}
