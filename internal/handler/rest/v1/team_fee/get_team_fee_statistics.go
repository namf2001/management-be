package team_fee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTeamFeeStatistics handles the request to get statistics for team fees
func (h Handler) GetTeamFeeStatistics(ctx *gin.Context) {
	// Get year filter from query parameter (optional)
	yearStr := ctx.Query("year")

	// Parse year if provided
	var year *int
	if yearStr != "" {
		yearInt, err := strconv.Atoi(yearStr)
		if err == nil {
			year = &yearInt
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid year format",
			})
			return
		}
	}

	// Get team fee statistics from controller
	stats, err := h.teamFeeCtrl.GetTeamFeeStatistics(ctx.Request.Context(), year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve team fee statistics",
		})
		return
	}

	// Prepare monthly summaries
	monthlySummaries := make([]MonthlyFeeSummaryResponse, len(stats.MonthlySummary))
	for i, month := range stats.MonthlySummary {
		monthlySummaries[i] = MonthlyFeeSummaryResponse{
			Month:            month.Month,
			TotalAmount:      month.TotalAmount,
			NumberOfPayments: month.NumberOfPayments,
		}
	}

	// Prepare yearly summaries
	yearlySummaries := make([]YearlyFeeSummaryResponse, len(stats.YearlySummary))
	for i, year := range stats.YearlySummary {
		yearlySummaries[i] = YearlyFeeSummaryResponse{
			Year:             year.Year,
			TotalAmount:      year.TotalAmount,
			NumberOfPayments: year.NumberOfPayments,
		}
	}

	// Return statistics
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": TeamFeeStatisticsResponse{
			Summary: TeamFeeSummaryResponse{
				TotalAmount:   stats.Summary.TotalAmount,
				TotalPayments: stats.Summary.TotalPayments,
				AverageAmount: stats.Summary.AverageAmount,
			},
			MonthlySummary: monthlySummaries,
			YearlySummary:  yearlySummaries,
		},
	})
}
