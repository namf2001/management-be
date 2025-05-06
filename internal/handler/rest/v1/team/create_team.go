package team

import (
	"management-be/internal/pkg/unit"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTeamRequest represents the request body for creating a team
// @name CreateTeamRequest
type CreateTeamRequest struct {
	Name          string `json:"name" validate:"required,min=2,max=100" example:"FC Barcelona"`
	CompanyName   string `json:"company_name" validate:"required" example:"FC Barcelona Sports Club"`
	ContactPerson string `json:"contact_person" validate:"required" example:"Joan Laporta"`
	ContactPhone  string `json:"contact_phone" validate:"omitempty,min=5,max=20" example:"123456789"`
	ContactEmail  string `json:"contact_email" validate:"omitempty,email" example:"contact@fcbarcelona.com"`
}

// CreateTeam
// @Summary      Create a new team
// @Description  Create a new team with name, company name, and contact information
// @Tags         teams
// @Accept       json
// @Produce      json
// @Param        team  body      CreateTeamRequest  true  "Team information"
// @Success      201  {object}  object{success=bool,data=TeamResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/teams [post]
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
