package team_fee

import (
	"context"
	"testing"
	"time"

	"management-be/internal/model"
	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

// TestCreateTeamFee tests the CreateTeamFee function
func TestCreateTeamFee(t *testing.T) {
	type args struct {
		input     CreateTeamFeeInput
		expResult model.TeamFee
		expErr    error
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	tcs := map[string]args{
		"success": {
			input: CreateTeamFeeInput{
				Amount:      100.50,
				PaymentDate: now,
				Description: "Monthly team fee",
			},
			expResult: model.TeamFee{
				Amount:      100.50,
				PaymentDate: now,
				Description: "Monthly team fee",
			},
		},
		"success with future date": {
			input: CreateTeamFeeInput{
				Amount:      200.75,
				PaymentDate: tomorrow,
				Description: "Advance payment",
			},
			expResult: model.TeamFee{
				Amount:      200.75,
				PaymentDate: tomorrow,
				Description: "Advance payment",
			},
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load team data first
				repo := NewRepository(tx.Client())
				teamFee, err := repo.CreateTeamFee(context.Background(), tc.input)

				// then
				if tc.expErr != nil {
					require.Error(t, err)
					require.Contains(t, err.Error(), tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotZero(t, teamFee.ID)
					require.Equal(t, tc.input.Amount, teamFee.Amount)
					// Compare dates by truncating to day precision to avoid timezone issues
					require.Equal(t, tc.input.PaymentDate.Format("2006-01-02"), teamFee.PaymentDate.Format("2006-01-02"))
					require.Equal(t, tc.input.Description, teamFee.Description)
					require.NotZero(t, teamFee.CreatedAt)
					require.NotZero(t, teamFee.UpdatedAt)
					require.Nil(t, teamFee.DeletedAt)

					// Verify the team fee was actually created in the database
					dbTeamFee, err := tx.Client().TeamFee.Get(context.Background(), teamFee.ID)
					require.NoError(t, err)
					require.Equal(t, tc.input.Amount, dbTeamFee.Amount)
					// Compare dates by truncating to day precision to avoid timezone issues
					require.Equal(t, tc.input.PaymentDate.Format("2006-01-02"), dbTeamFee.PaymentDate.Format("2006-01-02"))
					require.Equal(t, tc.input.Description, dbTeamFee.Description)
				}
			})
		})
	}
}
