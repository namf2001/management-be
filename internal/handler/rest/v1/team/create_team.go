package team

import (
	"management-be/internal/pkg/unit"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTeamRequest represents the request body for creating a team
type CreateTeamRequest struct {
	Name          string `json:"name" validate:"required,min=2,max=100"`
	CompanyName   string `json:"company_name" validate:"required"`
	ContactPerson string `json:"contact_person" validate:"required"`
	ContactPhone  string `json:"contact_phone" validate:"omitempty,min=5,max=20"`
	ContactEmail  string `json:"contact_email" validate:"omitempty,email"`
}

// CreateTeam handles the request to create a new team
func (h Handler) CreateTeam(ctx *gin.Context) {
	var req CreateTeamRequest
	if !unit.ValidateJSON(ctx, &req) {
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
