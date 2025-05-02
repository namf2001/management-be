package team_fee

import (
	"management-be/internal/controller/team_fee"
)

// Handler represents the team fee handler
type Handler struct {
	teamFeeCtrl team_fee.Controller
}

// NewHandler creates a new team fee handler
func NewHandler(teamFeeCtrl team_fee.Controller) Handler {
	return Handler{
		teamFeeCtrl: teamFeeCtrl,
	}
}

// TeamFeeResponse represents the response format for team fee data
type TeamFeeResponse struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
	Description string  `json:"description"`
}

// TeamFeeSummaryResponse represents the response format for team fee summary
type TeamFeeSummaryResponse struct {
	TotalAmount   float64 `json:"total_amount"`
	TotalPayments int     `json:"total_payments"`
	AverageAmount float64 `json:"average_amount"`
}

// TeamFeeListResponse represents the response format for listing team fees
type TeamFeeListResponse struct {
	TeamFees []TeamFeeResponse      `json:"team_fees"`
	Summary  TeamFeeSummaryResponse `json:"summary"`
}

// MonthlyFeeSummaryResponse represents monthly fee summary response
type MonthlyFeeSummaryResponse struct {
	Month            string  `json:"month"`
	TotalAmount      float64 `json:"total_amount"`
	NumberOfPayments int     `json:"number_of_payments"`
}

// YearlyFeeSummaryResponse represents yearly fee summary response
type YearlyFeeSummaryResponse struct {
	Year             int     `json:"year"`
	TotalAmount      float64 `json:"total_amount"`
	NumberOfPayments int     `json:"number_of_payments"`
}

// TeamFeeStatisticsResponse represents the response format for team fee statistics
type TeamFeeStatisticsResponse struct {
	Summary        TeamFeeSummaryResponse      `json:"summary"`
	MonthlySummary []MonthlyFeeSummaryResponse `json:"monthly_summary"`
	YearlySummary  []YearlyFeeSummaryResponse  `json:"yearly_summary"`
}
