package player

import (
	"github.com/gin-gonic/gin"
	v1 "management-be/internal/handler/rest/v1"
	"net/http"
	"strconv"
)

type DeletePlayerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// DeletePlayer handles the request to delete a player by ID
func (h Handler) DeletePlayer(ctx *gin.Context) {
	// Get player ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorInfo{
			Code:    http.StatusBadRequest,
			Message: "Invalid player ID",
		})
		return
	}

	// First check if player exists
	_, err = h.playerCtrl.GetPlayerByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, v1.ErrorInfo{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	// Delete player
	err = h.playerCtrl.DeletePlayer(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, v1.ErrorInfo{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, DeletePlayerResponse{
		Success: true,
		Message: "Player deleted successfully",
	})
}
