package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeletePlayer handles the request to delete a player by ID
func (h Handler) DeletePlayer(ctx *gin.Context) {
	// Get player ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid player ID",
		})
		return
	}

	// First check if player exists
	_, err = h.playerCtrl.GetPlayerByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Player not found",
		})
		return
	}

	// Delete player
	err = h.playerCtrl.DeletePlayer(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete player",
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Player deleted successfully",
	})
}
