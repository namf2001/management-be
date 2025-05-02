package team_fee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteTeamFee handles the request to delete a team fee
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
