package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/team_fee"
	"time"
)

// UpdateTeamFeeInput represents input for updating a team fee
type UpdateTeamFeeInput struct {
	Amount      float64
	PaymentDate time.Time
	Description string
}

// UpdateTeamFee updates an existing team fee
func (i impl) UpdateTeamFee(ctx context.Context, id int, input UpdateTeamFeeInput) (model.TeamFee, error) {
	// First check if the fee exists
	_, err := i.repo.TeamFee().GetTeamFeeByID(ctx, id)
	if err != nil {
		return model.TeamFee{}, err
	}

	return i.repo.TeamFee().UpdateTeamFee(ctx, id, team_fee.UpdateTeamFeeInput{
		Amount:      input.Amount,
		PaymentDate: input.PaymentDate,
		Description: input.Description,
	})
}
