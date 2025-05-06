package team_fee

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"
)

func TestDeleteTeamFee(t *testing.T) {
	type args struct {
		id     int
		expErr error
	}

	tcs := map[string]args{
		"success": {
			id:     1,
			expErr: nil,
		},
		"not_found": {
			id:     999,
			expErr: ErrNotFound,
		},
	}
	for _, tc := range tcs {
		testent.WithEntTx(t, func(tx *ent.Tx) {
			// Load team fee data first
			testent.LoadTestSQLFile(t, tx, "testdata/insert_team_fee.sql")
			// Load team data first
			repo := NewRepository(tx.Client())
			err := repo.DeleteTeamFee(context.Background(), tc.id)

			// then
			if tc.expErr != nil {
				require.ErrorIs(t, err, tc.expErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
