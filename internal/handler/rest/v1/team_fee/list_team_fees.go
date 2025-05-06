package team_fee

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ListTeamFees
// @Summary      List all team fees
// @Description  Get a list of all team fees with optional date range filters
// @Tags         team-fees
// @Accept       json
// @Produce      json
// @Param        start_date  query     string  false  "Start date (YYYY-MM-DD)"
// @Param        end_date    query     string  false  "End date (YYYY-MM-DD)"
// @Success      200  {object}  object{success=bool,data=TeamFeeListResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/team-fees [get]
func (h Handler) ListTeamFees(ctx *gin.Context) {
	// Get date range filters from query parameters
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	// Parse date filters if provided
	var startDate, endDate *time.Time
	if startDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			startDate = &parsedDate
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid start_date format. Use YYYY-MM-DD",
			})
			return
		}
	}

	if endDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			endDate = &parsedDate
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid end_date format. Use YYYY-MM-DD",
			})
			return
		}
	}

	// Get team fees with filters from controller
	teamFees, summary, err := h.teamFeeCtrl.ListTeamFees(ctx.Request.Context(), startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve team fees",
		})
		return
	}

	// Convert team fees to response format
	response := make([]TeamFeeResponse, len(teamFees))
	for i, fee := range teamFees {
		response[i] = TeamFeeResponse{
			ID:          fee.ID,
			Amount:      fee.Amount,
			PaymentDate: fee.PaymentDate.Format("2006-01-02T15:04:05Z07:00"),
			Description: fee.Description,
		}
	}

	// Create summary response
	summaryResponse := TeamFeeSummaryResponse{
		TotalAmount:   summary.TotalAmount,
		TotalPayments: summary.TotalPayments,
		AverageAmount: summary.AverageAmount,
	}

	// Return team fees with summary
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": TeamFeeListResponse{
			TeamFees: response,
			Summary:  summaryResponse,
		},
	})
}
