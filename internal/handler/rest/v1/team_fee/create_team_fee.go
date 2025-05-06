package team_fee

import (
	"management-be/internal/controller/team_fee"
	"management-be/internal/pkg/unit"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTeamFeeRequest represents the request body for creating a team fee
// @name CreateTeamFeeRequest
type CreateTeamFeeRequest struct {
	Amount      float64 `json:"amount" validate:"required,gt=0" example:"1000000"`
	PaymentDate string  `json:"payment_date" validate:"required,datetime=2006-01-02" example:"2024-06-01"`
	Description string  `json:"description" validate:"required,min=2,max=200" example:"Tournament registration fee"`
}

// CreateTeamFee
// @Summary      Create a new team fee
// @Description  Create a new team fee with amount, payment date and description
// @Tags         team-fees
// @Accept       json
// @Produce      json
// @Param        team_fee  body      CreateTeamFeeRequest  true  "Team fee information"
// @Success      201  {object}  object{success=bool,data=TeamFeeResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/team-fees [post]
func (h Handler) CreateTeamFee(ctx *gin.Context) {
	var req CreateTeamFeeRequest
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

	// Create team fee
	input := team_fee.CreateTeamFeeInput{
		Amount:      req.Amount,
		PaymentDate: paymentDate,
		Description: req.Description,
	}

	teamFee, err := h.teamFeeCtrl.CreateTeamFee(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create team fee",
		})
		return
	}

	// Return created team fee
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": TeamFeeResponse{
			ID:          teamFee.ID,
			Amount:      teamFee.Amount,
			PaymentDate: teamFee.PaymentDate.Format("2006-01-02T15:04:05Z07:00"),
			Description: teamFee.Description,
		},
	})
}
