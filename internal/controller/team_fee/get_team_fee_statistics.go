package team_fee

import (
	"context"
	"management-be/internal/model"
	"time"
)

// GetTeamFeeStatistics returns statistics about team fees
func (i impl) GetTeamFeeStatistics(ctx context.Context, year *int) (model.TeamFeeStatistics, error) {
	// Get all team fees first (applying year filter if provided)
	var startDate, endDate *time.Time

	// If year is provided, set date range for that year
	if year != nil {
		start := time.Date(*year, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(*year, 12, 31, 23, 59, 59, 0, time.UTC)
		startDate = &start
		endDate = &end
	}

	teamFees, err := i.repo.TeamFee().ListTeamFees(ctx, startDate, endDate)
	if err != nil {
		return model.TeamFeeStatistics{}, err
	}

	// Initialize the statistics result
	stats := model.TeamFeeStatistics{
		Summary: model.TeamFeeSummary{
			TotalPayments: len(teamFees),
		},
		MonthlySummary: make([]model.MonthlyFeeSummary, 0),
		YearlySummary:  make([]model.YearlyFeeSummary, 0),
	}

	// Maps to track monthly and yearly summaries
	monthlyMap := make(map[string]model.MonthlyFeeSummary)
	yearlyMap := make(map[int]model.YearlyFeeSummary)

	// Process each fee to build the statistics
	var totalAmount float64
	for _, fee := range teamFees {
		totalAmount += fee.Amount

		// Get year and month
		feeYear := fee.PaymentDate.Year()
		feeMonth := fee.PaymentDate.Format("January")

		// Update monthly summary
		monthKey := fee.PaymentDate.Format("2006-01")
		month, exists := monthlyMap[monthKey]
		if !exists {
			month = model.MonthlyFeeSummary{
				Month: feeMonth,
			}
		}
		month.TotalAmount += fee.Amount
		month.NumberOfPayments++
		monthlyMap[monthKey] = month

		// Update yearly summary
		year, exists := yearlyMap[feeYear]
		if !exists {
			year = model.YearlyFeeSummary{
				Year: feeYear,
			}
		}
		year.TotalAmount += fee.Amount
		year.NumberOfPayments++
		yearlyMap[feeYear] = year
	}

	// Set total amount in summary
	stats.Summary.TotalAmount = totalAmount

	// Calculate average amount if there are payments
	if len(teamFees) > 0 {
		stats.Summary.AverageAmount = totalAmount / float64(len(teamFees))
	}

	// Convert maps to slices for the final result
	for _, month := range monthlyMap {
		stats.MonthlySummary = append(stats.MonthlySummary, month)
	}

	for _, year := range yearlyMap {
		stats.YearlySummary = append(stats.YearlySummary, year)
	}

	return stats, nil
}
