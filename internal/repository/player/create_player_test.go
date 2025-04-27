package player

import (
	"context"
	"testing"
	"time"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestCreatePlayer(t *testing.T) {
	type args struct {
		input  InputPlayer
		expErr error
	}

	tcs := map[string]args{
		"success": {
			input: InputPlayer{
				DepartmentID: 1,
				FullName:     "New Player",
				Position:     "Midfielder",
				JerseyNumber: 777,
				DateOfBirth:  time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
				HeightCm:     175,
				WeightKg:     70,
				Phone:        "0987654321",
				Email:        "new@example.com",
				IsActive:     true,
			},
		},
		"err - invalid department": {
			input: InputPlayer{
				DepartmentID: 999,
				FullName:     "Invalid Player",
				Position:     "Forward",
				JerseyNumber: 9,
				DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				HeightCm:     180,
				WeightKg:     75,
				Phone:        "1234567890",
				Email:        "invalid@example.com",
				IsActive:     true,
			},
			expErr: ErrDatabase,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data to ensure department exists
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")

				repo := NewRepository(tx.Client())
				player, err := repo.CreatePlayer(context.Background(), tc.input)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotZero(t, player.ID)
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

					// Verify the player was actually created in the database
					dbPlayer, err := tx.Client().Player.Get(context.Background(), player.ID)
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
				}
			})
		})
	}
}
