package team_fee

// TeamFeeResponse represents the response body for team fee operations
// @name TeamFeeResponse
type TeamFeeResponse struct {
	ID          int     `json:"id" example:"1"`
	Amount      float64 `json:"amount" example:"1000000"`
	PaymentDate string  `json:"payment_date" example:"2024-06-01T00:00:00Z"`
	Description string  `json:"description" example:"Tournament registration fee"`
}

// TeamFeeSummaryResponse represents the summary information for team fees
// @name TeamFeeSummaryResponse
type TeamFeeSummaryResponse struct {
	TotalAmount   float64 `json:"total_amount" example:"5000000"`
	TotalPayments int     `json:"total_payments" example:"5"`
	AverageAmount float64 `json:"average_amount" example:"1000000"`
}

// TeamFeeListResponse represents the response body for listing team fees
// @name TeamFeeListResponse
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
