package match

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateMatchRequest represents the request body for creating a match
type CreateMatchRequest struct {
	OpponentTeamID int       `json:"opponent_team_id" binding:"required"`
	MatchDate      time.Time `json:"match_date" binding:"required"`
	Venue          string    `json:"venue" binding:"required"`
	IsHomeGame     bool      `json:"is_home_game"`
	Notes          string    `json:"notes"`
}

// CreateMatchResponse represents the response for creating a match
type CreateMatchResponse struct {
	ID             int       `json:"id"`
	OpponentTeamID int       `json:"opponent_team_id"`
	MatchDate      time.Time `json:"match_date"`
	Venue          string    `json:"venue"`
	IsHomeGame     bool      `json:"is_home_game"`
	Status         string    `json:"status"`
	Notes          string    `json:"notes"`
}

// CreateMatch handles the request to create a new match
func (h Handler) CreateMatch(ctx *gin.Context) {
	var req CreateMatchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Call the controller
	match, err := h.matchCtrl.CreateMatch(
		ctx.Request.Context(),
		req.OpponentTeamID,
		req.MatchDate,
		req.Venue,
		req.IsHomeGame,
		req.Notes,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create match",
		})
		return
	}

	// Prepare response
	response := CreateMatchResponse{
		ID:             match.ID,
		OpponentTeamID: match.OpponentTeamID,
		MatchDate:      match.MatchDate,
		Venue:          match.Venue,
		IsHomeGame:     match.IsHomeGame,
		Status:         match.Status,
		Notes:          match.Notes,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    response,
	})
}
