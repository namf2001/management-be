package team

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UpdateTeamRequest represents the request body for updating a team
type UpdateTeamRequest struct {
	Name          string `json:"name" binding:"required"`
	CompanyName   string `json:"company_name" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	ContactPhone  string `json:"contact_phone"`
	ContactEmail  string `json:"contact_email"`
}

// UpdateTeam handles the request to update an existing team
func (h Handler) UpdateTeam(ctx *gin.Context) {
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

	// Parse request body
	var req UpdateTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Update team in controller
	team, err := h.teamCtrl.UpdateTeam(
		ctx.Request.Context(),
		id,
		req.Name,
		req.CompanyName,
		req.ContactPerson,
		req.ContactPhone,
		req.ContactEmail,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update team",
		})
		return
	}

	// Return updated team
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": TeamResponse{
			ID:            team.ID,
			Name:          team.Name,
			CompanyName:   team.CompanyName,
			ContactPerson: team.ContactPerson,
			ContactPhone:  team.ContactPhone,
			ContactEmail:  team.ContactEmail,
		},
	})
}
