package team

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteTeam handles the request to delete a team
func (h Handler) DeleteTeam(ctx *gin.Context) {
	// Get team ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team ID",
		})
		return
	}

	// Delete team from controller
	err = h.teamCtrl.DeleteTeam(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete team",
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Team deleted successfully",
	})
}
