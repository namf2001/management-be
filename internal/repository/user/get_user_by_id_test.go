package user

import (
	"context"
	"errors"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/user"

	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	type args struct {
		userID int
		expErr error
	}

	tcs := map[string]args{
		"success": {
			userID: 0, // Will be set dynamically
		},
		"err - user not found": {
			userID: 999, // Non-existent user ID
			expErr: errors.New("user not found"),
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
				user, err := repo.GetUserByID(context.Background(), tc.userID)

				// then
				if tc.expErr != nil {
					require.Error(t, err)
					require.Contains(t, err.Error(), tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.userID, user.ID)
					require.Equal(t, "admintest", user.Username)
					require.Equal(t, "admintest@gmail.com", user.Email)
					require.Equal(t, "Admin Test User", user.FullName)
				}
			})
		})
	}
}
