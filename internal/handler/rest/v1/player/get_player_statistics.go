package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PositionStatsResponse struct {
	Position string `json:"position"`
	Count    int    `json:"count"`
}

type GoalsByMatchResponse struct {
	MatchDate string `json:"match_date"`
	Opponent  string `json:"opponent"`
	Goals     int    `json:"goals"`
}

type DetailedPlayerStatisticsResponse struct {
	TotalMatches       int32                  `json:"total_matches"`
	TotalMinutesPlayed int32                  `json:"total_minutes_played"`
	TotalGoals         int32                  `json:"total_goals"`
	TotalAssists       int32                  `json:"total_assists"`
	TotalYellowCards   int32                  `json:"total_yellow_cards"`
	TotalRedCards      int32                  `json:"total_red_cards"`
	MatchesByPosition  map[string]int32       `json:"matches_by_position"`
	GoalsByMatch       []GoalsByMatchResponse `json:"goals_by_match"`
}

// GetPlayerStatistics handles the request to get detailed statistics for a player
func (h Handler) GetPlayerStatistics(ctx *gin.Context) {
	// Get player ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid player ID",
		})
		return
	}

	// First check if player exists
	_, err = h.playerCtrl.GetPlayerByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Player not found",
		})
		return
	}

	// Get player statistics
	stats, err := h.playerCtrl.GetPlayerStatistics(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve player statistics",
		})
		return
	}

	// Format goals by match
	goalsByMatch := make([]GoalsByMatchResponse, 0) // Initialize empty slice since GoalsByMatch is not in the model
	matchesByPosition := make(map[string]int32)     // Initialize empty map since MatchesByPosition is not in the model

	// Return detailed statistics
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": DetailedPlayerStatisticsResponse{
			TotalMatches:       stats.TotalMatches,
			TotalMinutesPlayed: stats.TotalMinutesPlayed,
			TotalGoals:         stats.TotalGoals,
			TotalAssists:       stats.TotalAssists,
			TotalYellowCards:   stats.TotalYellowCards,
			TotalRedCards:      stats.TotalRedCards,
			MatchesByPosition:  matchesByPosition,
			GoalsByMatch:       goalsByMatch,
		},
	})
}
