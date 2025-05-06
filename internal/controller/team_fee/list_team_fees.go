package team_fee

import (
	"context"
	"management-be/internal/model"
	"time"
)

// ListTeamFees returns all team fees with optional filters
func (i impl) ListTeamFees(ctx context.Context, startDate, endDate *time.Time) ([]model.TeamFee, model.TeamFeeSummary, error) {
	teamFees, err := i.repo.TeamFee().ListTeamFees(ctx, startDate, endDate)
	if err != nil {
		return nil, model.TeamFeeSummary{}, err
	}

	// Calculate summary statistics
	var totalAmount float64
	for _, fee := range teamFees {
		totalAmount += fee.Amount
	}

	// Create summary
	summary := model.TeamFeeSummary{
		TotalAmount:   totalAmount,
		TotalPayments: len(teamFees),
	}

	// Calculate average if there are payments
	if len(teamFees) > 0 {
		summary.AverageAmount = totalAmount / float64(len(teamFees))
	}

	return teamFees, summary, nil
}
