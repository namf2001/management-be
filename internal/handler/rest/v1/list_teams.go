package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListTeams handles the request to list all teams with pagination
func (h Handler) ListTeams(ctx *gin.Context) {
	// Get pagination parameters from query string
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	// Get teams from controller
	teams, total, err := h.teamCtrl.GetAllTeams(ctx.Request.Context(), page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve teams",
		})
		return
	}

	// Convert teams to response format
	response := make([]TeamResponse, len(teams))
	for i, team := range teams {
		response[i] = TeamResponse{
			ID:            team.ID,
			Name:          team.Name,
			CompanyName:   team.CompanyName,
			ContactPerson: team.ContactPerson,
			ContactPhone:  team.ContactPhone,
			ContactEmail:  team.ContactEmail,
		}
	}

	// Calculate pagination info
	totalPages := (total + limit - 1) / limit // Ceiling division
	if totalPages < 1 {
		totalPages = 1
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"teams": response,
			"pagination": gin.H{
				"current_page":   page,
				"total_pages":    totalPages,
				"total_items":    total,
				"items_per_page": limit,
			},
		},
	})
}