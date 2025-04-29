package user

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/user"

	"github.com/stretchr/testify/require"
)

func TestUpdatePassword(t *testing.T) {
	type args struct {
		userID         int
		hashedPassword string
		expErr         error
	}

	tcs := map[string]args{
		"success": {
			userID:         0, // Will be set dynamically
			hashedPassword: "newpassword123",
		},
		"err - user not found": {
			userID:         999, // Non-existent user ID
			hashedPassword: "newpassword123",
			expErr:         ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_user.sql")

				// For the success case, get the user ID dynamically
				if s == "success" {
					user, err := tx.Client().User.Query().Where(user.UsernameEQ("admintest")).Only(context.Background())
					require.NoError(t, err)
					tc.userID = user.ID
				}

				repo := NewRepository(tx.Client())
				err := repo.UpdatePassword(context.Background(), tc.userID, tc.hashedPassword)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)

					// Verify the password was updated
					updatedUser, err := tx.Client().User.Get(context.Background(), tc.userID)
					require.NoError(t, err)
					require.Equal(t, tc.hashedPassword, updatedUser.Password)
				}
			})
		})
	}
}
