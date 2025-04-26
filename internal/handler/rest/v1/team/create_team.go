package team

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTeamRequest represents the request body for creating a team
type CreateTeamRequest struct {
	Name          string `json:"name" binding:"required"`
	CompanyName   string `json:"company_name" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	ContactPhone  string `json:"contact_phone"`
	ContactEmail  string `json:"contact_email"`
}

// CreateTeam handles the request to create a new team
func (h Handler) CreateTeam(ctx *gin.Context) {
	var req CreateTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	team, err := h.teamCtrl.CreateTeam(
		ctx.Request.Context(),
		req.Name,
		req.CompanyName,
		req.ContactPerson,
		req.ContactPhone,
		req.ContactEmail,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create team",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
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
