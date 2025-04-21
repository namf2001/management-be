package user

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		username string
		email    string
		password string
		fullName string
		expErr   error
	}

	tcs := map[string]args{
		"success": {
			username: "newuser",
			email:    "newuser@example.com",
			password: "password123",
			fullName: "New Test User",
		},
		"err - duplicate username": {
			username: "admintest", // Same as in insert_user.sql
			email:    "different@example.com",
			password: "password123",
			fullName: "Duplicate Username User",
			expErr:   ErrDatabase,
		},
		"err - duplicate email": {
			username: "differentuser",
			email:    "admintest@gmail.com", // Same as in insert_user.sql
			password: "password123",
			fullName: "Duplicate Email User",
			expErr:   ErrDatabase,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_user.sql")

				repo := NewRepository(tx.Client())
				user, err := repo.CreateUser(context.Background(), tc.username, tc.email, tc.password, tc.fullName)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotZero(t, user.ID)
					require.Equal(t, tc.username, user.Username)
					require.Equal(t, tc.email, user.Email)
					require.Equal(t, tc.fullName, user.FullName)

					// Verify the user was actually created in the database
					dbUser, err := tx.Client().User.Get(context.Background(), user.ID)
					require.NoError(t, err)
					require.Equal(t, tc.username, dbUser.Username)
					require.Equal(t, tc.email, dbUser.Email)
					require.Equal(t, tc.password, dbUser.Password)
					require.Equal(t, tc.fullName, dbUser.FullName)
				}
			})
		})
	}
}
