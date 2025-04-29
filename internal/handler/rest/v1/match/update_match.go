package match

import (
	"errors"
	"management-be/internal/controller/match"
	"management-be/internal/pkg/unit"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdateMatchRequest represents the request body for updating a match
type UpdateMatchRequest struct {
	OpponentTeamID int       `json:"opponent_team_id" validate:"required,min=1"`
	MatchDate      time.Time `json:"match_date" validate:"required"`
	Venue          string    `json:"venue" validate:"required,min=2"`
	IsHomeGame     bool      `json:"is_home_game"`
	OurScore       int32     `json:"our_score" validate:"min=0"`
	OpponentScore  int32     `json:"opponent_score" validate:"min=0"`
	Status         string    `json:"status" validate:"required,oneof=scheduled in_progress completed cancelled"`
	Notes          string    `json:"notes" validate:"omitempty,max=1000"`
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
	if !unit.ValidateJSON(ctx, &req) {
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
			"error":   err.Error(),
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
