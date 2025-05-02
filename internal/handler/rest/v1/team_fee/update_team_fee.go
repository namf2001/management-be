package team_fee

import (
	"management-be/internal/controller/team_fee"
	"management-be/internal/pkg/unit"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdateTeamFeeRequest represents the request body for updating a team fee
type UpdateTeamFeeRequest struct {
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	PaymentDate string  `json:"payment_date" validate:"required,datetime=2006-01-02"`
	Description string  `json:"description" validate:"required,min=2,max=200"`
}

// UpdateTeamFee handles the request to update an existing team fee
func (h Handler) UpdateTeamFee(ctx *gin.Context) {
	// Get team fee ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team fee ID",
		})
		return
	}

	var req UpdateTeamFeeRequest
	if !unit.ValidateJSON(ctx, &req) {
		return
	}

	// Parse payment date
	paymentDate, err := time.Parse("2006-01-02", req.PaymentDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid payment date format. Use YYYY-MM-DD",
		})
		return
	}

	// Update team fee
	input := team_fee.UpdateTeamFeeInput{
		Amount:      req.Amount,
		PaymentDate: paymentDate,
		Description: req.Description,
	}

	teamFee, err := h.teamFeeCtrl.UpdateTeamFee(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update team fee",
		})
		return
	}

	// Return updated team fee
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": TeamFeeResponse{
			ID:          teamFee.ID,
			Amount:      teamFee.Amount,
			PaymentDate: teamFee.PaymentDate.Format("2006-01-02T15:04:05Z07:00"),
			Description: teamFee.Description,
		},
	})
}
