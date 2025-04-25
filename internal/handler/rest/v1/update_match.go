package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"management-be/internal/controller/match"
	"net/http"
	"strconv"
	"time"
)

// UpdateMatchRequest represents the request body for updating a match
type UpdateMatchRequest struct {
	OpponentTeamID int       `json:"opponent_team_id" binding:"required"`
	MatchDate      time.Time `json:"match_date" binding:"required"`
	Venue          string    `json:"venue" binding:"required"`
	IsHomeGame     bool      `json:"is_home_game"`
	OurScore       int32     `json:"our_score"`
	OpponentScore  int32     `json:"opponent_score"`
	Status         string    `json:"status" binding:"required"`
	Notes          string    `json:"notes"`
}

// UpdateMatchResponse represents the response for updating a match
type UpdateMatchResponse struct {
	ID             int       `json:"id"`
	OpponentTeamID int       `json:"opponent_team_id"`
	MatchDate      time.Time `json:"match_date"`
	Venue          string    `json:"venue"`
	IsHomeGame     bool      `json:"is_home_game"`
	OurScore       int32     `json:"our_score"`
	OpponentScore  int32     `json:"opponent_score"`
	Status         string    `json:"status"`
	Notes          string    `json:"notes"`
}

// UpdateMatch handles the request to update an existing match
func (h Handler) UpdateMatch(ctx *gin.Context) {
	// Parse match ID from URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid match ID",
		})
		return
	}

	var req UpdateMatchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Call the controller
	matchUpdate, err := h.matchCtrl.UpdateMatch(
		ctx.Request.Context(),
		id,
		req.OpponentTeamID,
		req.MatchDate,
		req.Venue,
		req.IsHomeGame,
		req.OurScore,
		req.OpponentScore,
		req.Status,
		req.Notes,
	)
	if err != nil {
		if errors.Is(err, match.ErrMatchNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Match not found",
			})
			return
		}
		// Fallback to 500 for all other errors
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update match",
		})
		return
	}

	// Prepare response
	response := UpdateMatchResponse{
		ID:             matchUpdate.ID,
		OpponentTeamID: matchUpdate.OpponentTeamID,
		MatchDate:      matchUpdate.MatchDate,
		Venue:          matchUpdate.Venue,
		IsHomeGame:     matchUpdate.IsHomeGame,
		OurScore:       matchUpdate.OurScore,
		OpponentScore:  matchUpdate.OpponentScore,
		Status:         matchUpdate.Status,
		Notes:          matchUpdate.Notes,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
