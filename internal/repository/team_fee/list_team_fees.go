package team_fee

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent/teamfee"
	"time"
)

// ListTeamFees returns all team fees with optional filters
func (i impl) ListTeamFees(ctx context.Context, startDate, endDate *time.Time) ([]model.TeamFee, error) {
	query := i.entClient.TeamFee.Query().
		Where(teamfee.DeletedAtIsNil())

	// Apply date range filters if provided
	if startDate != nil {
		query = query.Where(teamfee.PaymentDateGTE(*startDate))
	}
	if endDate != nil {
		query = query.Where(teamfee.PaymentDateLTE(*endDate))
	}

	// Order by payment date (most recent first)
	query = query.Order(teamfee.ByPaymentDate())

	teamFees, err := query.All(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	// Convert to model format
	result := make([]model.TeamFee, len(teamFees))
	for i, fee := range teamFees {
		result[i] = model.TeamFee{
			ID:          fee.ID,
			Amount:      fee.Amount,
			PaymentDate: fee.PaymentDate,
			Description: fee.Description,
			CreatedAt:   fee.CreatedAt,
			UpdatedAt:   fee.UpdatedAt,
		}
	}

	return result, nil
}
