package team

import (
	"management-be/internal/pkg/unit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateTeamRequest represents the request body for updating a team
// @name UpdateTeamRequest
type UpdateTeamRequest struct {
	Name          string `json:"name" validate:"required,min=2,max=100" example:"FC Barcelona Updated"`
	CompanyName   string `json:"company_name" validate:"required" example:"FC Barcelona Sports Club Updated"`
	ContactPerson string `json:"contact_person" validate:"required" example:"Updated Contact Person"`
	ContactPhone  string `json:"contact_phone" validate:"omitempty,min=5,max=20" example:"987654321"`
	ContactEmail  string `json:"contact_email" validate:"omitempty,email" example:"updated@fcbarcelona.com"`
}

// UpdateTeam
// @Summary      Update a team
// @Description  Update an existing team's information
// @Tags         teams
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Team ID"
// @Param        team  body      UpdateTeamRequest  true  "Team information"
// @Success      200  {object}  object{success=bool,data=TeamResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/teams/{id} [put]
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
	if !unit.ValidateJSON(ctx, &req) {
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
