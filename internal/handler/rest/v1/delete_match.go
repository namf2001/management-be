package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteMatch handles the request to delete a match by ID
func (h Handler) DeleteMatch(ctx *gin.Context) {
	// Parse match ID from URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid match ID",
		})
		return
	}

	// Call the controller
	err = h.matchCtrl.DeleteMatch(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete match",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Match deleted successfully",
	})
}
