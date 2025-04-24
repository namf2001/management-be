package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ListMatchesResponse represents the response for listing matches
type ListMatchesResponse struct {
	ID            int          `json:"id"`
	OpponentTeam  TeamResponse `json:"opponent_team"`
	MatchDate     time.Time    `json:"match_date"`
	Venue         string       `json:"venue"`
	IsHomeGame    bool         `json:"is_home_game"`
	OurScore      int32        `json:"our_score,omitempty"`
	OpponentScore int32        `json:"opponent_score,omitempty"`
	Status        string       `json:"status"`
	Notes         string       `json:"notes"`
}

// ListMatches handles the request to list all matches with optional filters
func (h Handler) ListMatches(ctx *gin.Context) {
	// Parse query parameters
	status := ctx.Query("status")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")
	opponentTeamIDStr := ctx.Query("opponent_team_id")

	var startDate, endDate time.Time
	var opponentTeamID int

	// Parse start_date if provided
	if startDateStr != "" {
		var err error
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid start_date format",
			})
			return
		}
	}

	// Parse end_date if provided
	if endDateStr != "" {
		var err error
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid end_date format",
			})
			return
		}
	}

	// Parse opponent_team_id if provided
	if opponentTeamIDStr != "" {
		var err error
		_, err = fmt.Sscanf(opponentTeamIDStr, "%d", &opponentTeamID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid opponent_team_id",
			})
			return
		}
	}

	// Call the controller
	matches, err := h.matchCtrl.ListMatches(
		ctx.Request.Context(),
		status,
		startDate,
		endDate,
		opponentTeamID,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to list matches",
		})
		return
	}

	// Prepare response
	var response []ListMatchesResponse
	for _, match := range matches {
		response = append(response, ListMatchesResponse{
			ID: match.ID,
			OpponentTeam: TeamResponse{
				ID:          match.OpponentTeamID,
				Name:        "", // Need to fetch team details
				CompanyName: "",
			},
			MatchDate:     match.MatchDate,
			Venue:         match.Venue,
			IsHomeGame:    match.IsHomeGame,
			OurScore:      match.OurScore,
			OpponentScore: match.OpponentScore,
			Status:        match.Status,
			Notes:         match.Notes,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"matches": response,
		},
	})
}
