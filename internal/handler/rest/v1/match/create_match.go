package match

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"management-be/internal/controller/match"
	"management-be/internal/pkg/unit"
)

// CreateMatchRequest represents the request body for creating a match
type CreateMatchRequest struct {
	OpponentTeamID int       `json:"opponent_team_id" validate:"required,min=1"`
	MatchDate      time.Time `json:"match_date" validate:"required"`
	Venue          string    `json:"venue" validate:"required,min=2"`
	IsHomeGame     bool      `json:"is_home_game"`
	Notes          string    `json:"notes" validate:"omitempty,max=1000"`
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
	if !unit.ValidateJSON(ctx, &req) {
		return
	}

	// Call the controller
	match, err := h.matchCtrl.CreateMatch(
		ctx.Request.Context(),
		match.CreateMatchInput{
			OpponentTeamID: req.OpponentTeamID,
			MatchDate:      req.MatchDate,
			Venue:          req.Venue,
			IsHomeGame:     req.IsHomeGame,
			Notes:          req.Notes,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": CreateMatchResponse{
			ID:             match.ID,
			OpponentTeamID: match.OpponentTeamID,
			MatchDate:      match.MatchDate,
			Venue:          match.Venue,
			IsHomeGame:     match.IsHomeGame,
			Status:         match.Status,
			Notes:          match.Notes,
		},
	})
}
