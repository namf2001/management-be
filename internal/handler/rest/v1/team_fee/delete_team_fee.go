package team_fee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteTeamFee
// @Summary      Delete a team fee
// @Description  Delete an existing team fee by ID
// @Tags         team-fees
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Team Fee ID"
// @Success      200  {object}  object{success=bool,message=string}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/team-fees/{id} [delete]
func (h Handler) DeleteTeamFee(ctx *gin.Context) {
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

	// Delete team fee
	err = h.teamFeeCtrl.DeleteTeamFee(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete team fee",
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Team fee deleted successfully",
	})
}
