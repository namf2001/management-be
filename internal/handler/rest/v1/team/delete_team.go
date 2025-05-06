package team

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteTeam
// @Summary      Delete a team
// @Description  Delete an existing team by ID
// @Tags         teams
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Team ID"
// @Success      200  {object}  object{success=bool,message=string}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/teams/{id} [delete]
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
