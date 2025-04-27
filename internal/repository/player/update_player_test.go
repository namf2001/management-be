package player

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestUpdatePlayer(t *testing.T) {
	type args struct {
		id     int
		input  InputPlayer
		expErr error
	}

	tcs := map[string]args{
		"success": {
			id: 1,
			input: InputPlayer{
				DepartmentID: 1,
				FullName:     "Updated Player",
				Position:     "Defender",
				JerseyNumber: 5,
				DateOfBirth:  time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC),
				HeightCm:     185,
				WeightKg:     80,
				Phone:        "9876543210",
				Email:        "updated@example.com",
				IsActive:     false,
			},
		},
		"err - not found": {
			id: 999,
			input: InputPlayer{
				DepartmentID: 1,
				FullName:     "Non-existent Player",
				Position:     "Forward",
				JerseyNumber: 9,
				DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				HeightCm:     180,
				WeightKg:     75,
				Phone:        "1234567890",
				Email:        "nonexistent@example.com",
				IsActive:     true,
			},
			expErr: ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data to ensure department and player exist
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

				// Record the time before update to verify updated_at is changed
				beforeUpdate := time.Now()
				time.Sleep(10 * time.Millisecond) // Ensure time difference

				repo := NewRepository(tx.Client())
				player, err := repo.UpdatePlayer(context.Background(), tc.id, tc.input)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.id, player.ID)
					require.Equal(t, tc.input.DepartmentID, player.DepartmentID)
					require.Equal(t, tc.input.FullName, player.FullName)
					require.Equal(t, tc.input.Position, player.Position)
					require.Equal(t, tc.input.JerseyNumber, *player.JerseyNumber)
					require.Equal(t, tc.input.DateOfBirth.UTC(), player.DateOfBirth.UTC())
					require.Equal(t, tc.input.HeightCm, *player.HeightCM)
					require.Equal(t, tc.input.WeightKg, *player.WeightKG)
					require.Equal(t, tc.input.Phone, player.Phone)
					require.Equal(t, tc.input.Email, player.Email)
					require.Equal(t, tc.input.IsActive, player.IsActive)
					require.NotZero(t, player.CreatedAt)
					require.NotZero(t, player.UpdatedAt)
					require.True(t, player.UpdatedAt.After(beforeUpdate), "UpdatedAt should be after the time before update")

					// Verify the player was actually updated in the database
					dbPlayer, err := tx.Client().Player.Get(context.Background(), tc.id)
					require.NoError(t, err)
					require.Equal(t, tc.input.DepartmentID, dbPlayer.DepartmentID)
					require.Equal(t, tc.input.FullName, dbPlayer.FullName)
					require.Equal(t, tc.input.Position, dbPlayer.Position)
					require.Equal(t, tc.input.JerseyNumber, dbPlayer.JerseyNumber)
					require.Equal(t, tc.input.DateOfBirth.UTC(), dbPlayer.DateOfBirth.UTC())
					require.Equal(t, tc.input.HeightCm, dbPlayer.HeightCm)
					require.Equal(t, tc.input.WeightKg, dbPlayer.WeightKg)
					require.Equal(t, tc.input.Phone, dbPlayer.Phone)
					require.Equal(t, tc.input.Email, dbPlayer.Email)
					require.Equal(t, tc.input.IsActive, dbPlayer.IsActive)
					require.NotZero(t, dbPlayer.CreatedAt)
					require.NotZero(t, dbPlayer.UpdatedAt)
					require.True(t, dbPlayer.UpdatedAt.After(beforeUpdate), "UpdatedAt in DB should be after the time before update")
				}
			})
		})
	}
}
