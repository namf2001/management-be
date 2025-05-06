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
// @name UpdateTeamFeeRequest
type UpdateTeamFeeRequest struct {
	Amount      float64 `json:"amount" validate:"required,gt=0" example:"2000000"`
	PaymentDate string  `json:"payment_date" validate:"required,datetime=2006-01-02" example:"2024-06-10"`
	Description string  `json:"description" validate:"required,min=2,max=200" example:"League membership fee"`
}

// UpdateTeamFee
// @Summary      Update a team fee
// @Description  Update an existing team fee's information
// @Tags         team-fees
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Team Fee ID"
// @Param        team_fee  body      UpdateTeamFeeRequest  true  "Team fee information"
// @Success      200  {object}  object{success=bool,data=TeamFeeResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/team-fees/{id} [put]
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
