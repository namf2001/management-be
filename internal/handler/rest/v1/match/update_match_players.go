package match

import (
	"errors"
	"github.com/gin-gonic/gin"
	"management-be/internal/controller/match"
	"management-be/internal/model"
	"net/http"
	"strconv"
)

// UpdateMatchPlayersRequest represents the request body for updating match players
type UpdateMatchPlayersRequest struct {
	Players []MatchPlayerUpdateRequest `json:"players" binding:"required"`
}

// MatchPlayerUpdateRequest represents a player's performance in a match for update
type MatchPlayerUpdateRequest struct {
	PlayerID      int  `json:"player_id" binding:"required"`
	MinutesPlayed int32  `json:"minutes_played"`
	GoalsScored   int32  `json:"goals_scored"`
	Assists       int32  `json:"assists"`
	YellowCards   int32  `json:"yellow_cards"`
	RedCard       bool   `json:"red_card"`
}

// UpdateMatchPlayersResponse represents the response for updating match players
type UpdateMatchPlayersResponse struct {
	MatchID int                   `json:"match_id"`
	Players []MatchPlayerResponse `json:"players"`
}

// UpdateMatchPlayers handles the request to update player participation in a match
func (h Handler) UpdateMatchPlayers(ctx *gin.Context) {
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

	var req UpdateMatchPlayersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Convert request to model.MatchPlayer
	matchPlayers := make([]model.MatchPlayer, len(req.Players))
	for i, player := range req.Players {
		matchPlayers[i] = model.MatchPlayer{
			MatchID:       id,
			PlayerID:      player.PlayerID,
			MinutesPlayed: player.MinutesPlayed,
			GoalsScored:   player.GoalsScored,
			Assists:       player.Assists,
			YellowCards:   player.YellowCards,
			RedCard:       player.RedCard,
		}
	}

	// Call the controller
	err = h.matchCtrl.UpdateMatchPlayers(ctx.Request.Context(), id, matchPlayers)
	if err != nil {
		// Handle specific error types
		switch {
		case errors.Is(err, match.ErrMatchNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Match not found",
			})
		case errors.Is(err, match.ErrInvalidPlayerData):
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid player data",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update match players",
			})
		}
		return
	}

	// Prepare response
	// Note: In a real implementation, you would fetch the updated match players with player names
	var playerResponses []MatchPlayerResponse
	for _, player := range req.Players {
		playerResponses = append(playerResponses, MatchPlayerResponse{
			PlayerID:      player.PlayerID,
			PlayerName:    "", // Need to fetch player name
			JerseyNumber:  0,  // Need to fetch jersey number
			Position:      "", // Need to fetch position
			MinutesPlayed: int(player.MinutesPlayed),
			GoalsScored:   int(player.GoalsScored),
			Assists:       int(player.Assists),
			YellowCards:   int(player.YellowCards),
			RedCard:       player.RedCard,
		})
	}

	response := UpdateMatchPlayersResponse{
		MatchID: id,
		Players: playerResponses,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
