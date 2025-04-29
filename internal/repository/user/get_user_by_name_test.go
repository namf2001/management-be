package user

import (
	"context"
	"testing"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestGetUserByUsername(t *testing.T) {
	type args struct {
		givenUsername string
		expErr        error
	}

	tcs := map[string]args{
		"success": {
			givenUsername: "admintest",
		},
		"err - user not found": {
			givenUsername: "nonexistentuser",
			expErr:        ErrNotFound,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				testent.LoadTestSQLFile(t, tx, "testdata/insert_user.sql")
				repo := NewRepository(tx.Client())
				user, err := repo.GetUserByUsername(context.Background(), tc.givenUsername)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.givenUsername, user.Username)
				}
			})
		})
	}
}
