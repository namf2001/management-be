package player

import (
	"github.com/gin-gonic/gin"
	v1 "management-be/internal/handler/rest/v1"
	"net/http"
	"strconv"
)

// DeletePlayerResponse represents the response for deleting a player
// @name DeletePlayerResponse
type DeletePlayerResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Player deleted successfully"`
}

// DeletePlayer
// @Summary      Delete a player
// @Description  Delete a player by their ID
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Player ID"
// @Success      200  {object}  DeletePlayerResponse
// @Failure      400  {object}  object{code=int,message=string}
// @Failure      404  {object}  object{code=int,message=string}
// @Failure      500  {object}  object{code=int,message=string}
// @Router       /api/players/{id} [delete]
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
