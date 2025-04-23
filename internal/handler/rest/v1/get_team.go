package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTeam handles the request to get a team by ID
func (h Handler) GetTeam(ctx *gin.Context) {
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

	// Get team from controller
	team, err := h.teamCtrl.GetTeamByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Team not found",
		})
		return
	}

	// Get team statistics
	stats, err := h.teamCtrl.GetTeamStatistics(ctx.Request.Context(), id)
	if err != nil {
		// If we can't get statistics, just return the team without them
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
		return
	}

	// Convert match dates to string format
	matches := make([]MatchResponse, len(stats.Matches))
	for i, match := range stats.Matches {
		matches[i] = MatchResponse{
			MatchID:       match.ID,
			MatchDate:     match.MatchDate.Format("2006-01-02T15:04:05Z07:00"),
			Venue:         match.Venue,
			OurScore:      match.OurScore,
			OpponentScore: match.OpponentScore,
			Status:        match.Status,
		}
	}

	// Return team with statistics
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": TeamWithStatsResponse{
			ID:            team.ID,
			Name:          team.Name,
			CompanyName:   team.CompanyName,
			ContactPerson: team.ContactPerson,
			ContactPhone:  team.ContactPhone,
			ContactEmail:  team.ContactEmail,
			MatchHistory: MatchHistoryResponse{
				TotalMatches: stats.TotalMatches,
				Wins:         stats.Wins,
				Losses:       stats.Losses,
				Draws:        stats.Draws,
				Matches:      matches,
			},
		},
	})
}
