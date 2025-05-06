package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/team_fee"
	"time"
)

// CreateTeamFeeInput represents input for creating a team fee
type CreateTeamFeeInput struct {
	Amount      float64
	PaymentDate time.Time
	Description string
}

// CreateTeamFee creates a new team fee
func (i impl) CreateTeamFee(ctx context.Context, input CreateTeamFeeInput) (model.TeamFee, error) {
	return i.repo.TeamFee().CreateTeamFee(ctx,
		team_fee.CreateTeamFeeInput{
			Amount:      input.Amount,
			PaymentDate: input.PaymentDate,
			Description: input.Description,
		})
}
