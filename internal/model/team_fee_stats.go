package model

// TeamFeeSummary represents a summary of team fee information
type TeamFeeSummary struct {
	TotalAmount   float64 `json:"total_amount"`
	TotalPayments int     `json:"total_payments"`
	AverageAmount float64 `json:"average_amount"`
}

// MonthlyFeeSummary represents monthly fee summary
type MonthlyFeeSummary struct {
	Month            string  `json:"month"`
	TotalAmount      float64 `json:"total_amount"`
	NumberOfPayments int     `json:"number_of_payments"`
}

// YearlyFeeSummary represents yearly fee summary
type YearlyFeeSummary struct {
	Year             int     `json:"year"`
	TotalAmount      float64 `json:"total_amount"`
	NumberOfPayments int     `json:"number_of_payments"`
}

// TeamFeeStatistics represents comprehensive statistics about team fees
type TeamFeeStatistics struct {
	Summary        TeamFeeSummary      `json:"summary"`
	MonthlySummary []MonthlyFeeSummary `json:"monthly_summary"`
	YearlySummary  []YearlyFeeSummary  `json:"yearly_summary"`
}
