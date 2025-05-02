package team_fee

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// UpdateTeamFeeInput represents the input for updating a team fee
type UpdateTeamFeeInput struct {
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Description string    `json:"description"`
}

// UpdateTeamFee updates an existing team fee
func (i impl) UpdateTeamFee(ctx context.Context, id int, teamFee UpdateTeamFeeInput) (model.TeamFee, error) {
	// Update the team fee in database
	updatedFee, err := i.entClient.TeamFee.UpdateOneID(id).
		SetAmount(teamFee.Amount).
		SetPaymentDate(teamFee.PaymentDate).
		SetDescription(teamFee.Description).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return model.TeamFee{}, pkgerrors.WithStack(err)
	}

	return model.TeamFee{
		ID:          updatedFee.ID,
		Amount:      updatedFee.Amount,
		PaymentDate: updatedFee.PaymentDate,
		Description: updatedFee.Description,
		CreatedAt:   updatedFee.CreatedAt,
		UpdatedAt:   updatedFee.UpdatedAt,
	}, nil
}
