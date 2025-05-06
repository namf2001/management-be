package team_fee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTeamFee
// @Summary      Get a team fee by ID
// @Description  Get detailed information about a specific team fee
// @Tags         team-fees
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Team Fee ID"
// @Success      200  {object}  object{success=bool,data=TeamFeeResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      404  {object}  object{success=bool,error=string}
// @Router       /api/team-fees/{id} [get]
func (h Handler) GetTeamFee(ctx *gin.Context) {
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

	// Get team fee from controller
	teamFee, err := h.teamFeeCtrl.GetTeamFee(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Team fee not found",
		})
		return
	}

	// Return team fee
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
