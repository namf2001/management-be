package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"management-be/internal/controller/match"
	"net/http"
	"strconv"
)

// MatchSummaryResponse represents the match summary statistics
type MatchSummaryResponse struct {
	TotalPlayers       int32 `json:"total_players"`
	TotalMinutesPlayed int32 `json:"total_minutes_played"`
	TotalGoals         int32 `json:"total_goals"`
	TotalAssists       int32 `json:"total_assists"`
	TotalYellowCards   int32 `json:"total_yellow_cards"`
	TotalRedCards      int32 `json:"total_red_cards"`
}

// PlayerPerformanceResponse represents a player's performance in a match for statistics
type PlayerPerformanceResponse struct {
	PlayerID      int    `json:"player_id"`
	PlayerName    string `json:"player_name"`
	Position      string `json:"position"`
	MinutesPlayed int32  `json:"minutes_played"`
	GoalsScored   int32  `json:"goals_scored"`
	Assists       int32  `json:"assists"`
	YellowCards   int32  `json:"yellow_cards"`
	RedCard       bool   `json:"red_card"`
}

// MatchStatisticsResponse represents the response for getting match statistics
type MatchStatisticsResponse struct {
	MatchSummary      MatchSummaryResponse        `json:"match_summary"`
	PlayerPerformance []PlayerPerformanceResponse `json:"player_performance"`
	PositionSummary   map[string]int              `json:"position_summary"`
}

// GetMatchStatistics handles the request to get match statistics
func (h Handler) GetMatchStatistics(ctx *gin.Context) {
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

	// Call the controller
	stats, err := h.matchCtrl.GetMatchStatistics(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, match.ErrMatchNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Match not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get match statistics",
		})
		return
	}

	// Prepare response
	playerPerformance := make([]PlayerPerformanceResponse, len(stats.PlayerPerformance))
	positionSummary := make(map[string]int)

	for i, player := range stats.PlayerPerformance {
		playerPerformance[i] = PlayerPerformanceResponse{
			PlayerID:      player.PlayerID,
			PlayerName:    player.PlayerName,
			Position:      player.Position,
			MinutesPlayed: player.MinutesPlayed,
			GoalsScored:   player.GoalsScored,
			Assists:       player.Assists,
			YellowCards:   player.YellowCards,
			RedCard:       player.RedCard,
		}

		// Count players by position
		positionSummary[player.Position]++
	}

	response := MatchStatisticsResponse{
		MatchSummary: MatchSummaryResponse{
			TotalPlayers:       stats.MatchSummary.TotalPlayers,
			TotalMinutesPlayed: stats.MatchSummary.TotalMinutesPlayed,
			TotalGoals:         stats.MatchSummary.TotalGoals,
			TotalAssists:       stats.MatchSummary.TotalAssists,
			TotalYellowCards:   stats.MatchSummary.TotalYellowCards,
			TotalRedCards:      stats.MatchSummary.TotalRedCards,
		},
		PlayerPerformance: playerPerformance,
		PositionSummary:   positionSummary,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
