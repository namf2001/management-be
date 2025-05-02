package team_fee

import (
	"context"
	"management-be/internal/model"
	"time"

	pkgerrors "github.com/pkg/errors"
)

// CreateTeamFeeInput represents the input for creating a team fee
type CreateTeamFeeInput struct {
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Description string    `json:"description"`
}

// CreateTeamFee creates a new team fee
func (i impl) CreateTeamFee(ctx context.Context, teamFee CreateTeamFeeInput) (model.TeamFee, error) {
	// Create new team fee in database
	createdFee, err := i.entClient.TeamFee.Create().
		SetAmount(teamFee.Amount).
		SetPaymentDate(teamFee.PaymentDate).
		SetDescription(teamFee.Description).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return model.TeamFee{}, pkgerrors.WithStack(err)
	}

	// Return created team fee
	return model.TeamFee{
		ID:          createdFee.ID,
		Amount:      createdFee.Amount,
		PaymentDate: createdFee.PaymentDate,
		Description: createdFee.Description,
		CreatedAt:   createdFee.CreatedAt,
		UpdatedAt:   createdFee.UpdatedAt,
	}, nil
}
