package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerResponse struct {
	ID             int    `json:"id"`
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
	FullName       string `json:"full_name"`
	JerseyNumber   int    `json:"jersey_number"`
	Position       string `json:"position"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	HeightCm       int    `json:"height_cm,omitempty"`
	WeightKg       int    `json:"weight_kg,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	IsActive       bool   `json:"is_active"`
}

type PlayerStatisticsResponse struct {
	TotalMatches       int `json:"total_matches"`
	TotalMinutesPlayed int `json:"total_minutes_played"`
	TotalGoals         int `json:"total_goals"`
	TotalAssists       int `json:"total_assists"`
	TotalYellowCards   int `json:"total_yellow_cards"`
	TotalRedCards      int `json:"total_red_cards"`
}

type PlayerRecentMatchResponse struct {
	MatchID       int    `json:"match_id"`
	MatchDate     string `json:"match_date"`
	Opponent      string `json:"opponent"`
	MinutesPlayed int    `json:"minutes_played"`
	GoalsScored   int    `json:"goals_scored"`
	Assists       int    `json:"assists"`
}

type PlayerWithStatsResponse struct {
	PlayerResponse
	Statistics    PlayerStatisticsResponse    `json:"statistics"`
	RecentMatches []PlayerRecentMatchResponse `json:"recent_matches"`
}

// GetPlayer handles the request to get a player by ID
func (h Handler) GetPlayer(ctx *gin.Context) {
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

	// Get player from controller
	player, err := h.playerCtrl.GetPlayerByID(ctx.Request.Context(), id)
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
		// If we can't get statistics, just return the player without them
		var dateOfBirth string
		if player.DateOfBirth != nil && !player.DateOfBirth.IsZero() {
			dateOfBirth = player.DateOfBirth.Format("2006-01-02")
		}

		// Get jersey number, height and weight safely
		var jerseyNumber, heightCm, weightKg int
		if player.JerseyNumber != nil {
			jerseyNumber = int(*player.JerseyNumber)
		}
		if player.HeightCM != nil {
			heightCm = int(*player.HeightCM)
		}
		if player.WeightKG != nil {
			weightKg = int(*player.WeightKG)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": PlayerResponse{
				ID:             player.ID,
				DepartmentID:   player.DepartmentID,
				DepartmentName: "", // Department name would be fetched in a real implementation
				FullName:       player.FullName,
				JerseyNumber:   jerseyNumber,
				Position:       player.Position,
				DateOfBirth:    dateOfBirth,
				HeightCm:       heightCm,
				WeightKg:       weightKg,
				Phone:          player.Phone,
				Email:          player.Email,
				IsActive:       player.IsActive,
			},
		})
		return
	}

	// In a real implementation, you would fetch recent matches from a service
	// For now, create empty recent matches slice
	recentMatches := []PlayerRecentMatchResponse{}
	// If your API provides recent matches, you would populate this here
	// For example:
	// recentMatches := make([]PlayerRecentMatchResponse, len(matchData))
	// for i, match := range matchData {
	//     recentMatches[i] = PlayerRecentMatchResponse{...}
	// }

	// Format date of birth if it exists
	var dateOfBirth string
	if player.DateOfBirth != nil && !player.DateOfBirth.IsZero() {
		dateOfBirth = player.DateOfBirth.Format("2006-01-02")
	}

	// Get jersey number, height and weight safely
	var jerseyNumber, heightCm, weightKg int
	if player.JerseyNumber != nil {
		jerseyNumber = int(*player.JerseyNumber)
	}
	if player.HeightCM != nil {
		heightCm = int(*player.HeightCM)
	}
	if player.WeightKG != nil {
		weightKg = int(*player.WeightKG)
	}

	// Return player with statistics
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": PlayerWithStatsResponse{
			PlayerResponse: PlayerResponse{
				ID:             player.ID,
				DepartmentID:   player.DepartmentID,
				DepartmentName: "", // Department name would be fetched in a real implementation
				FullName:       player.FullName,
				JerseyNumber:   jerseyNumber,
				Position:       player.Position,
				DateOfBirth:    dateOfBirth,
				HeightCm:       heightCm,
				WeightKg:       weightKg,
				Phone:          player.Phone,
				Email:          player.Email,
				IsActive:       player.IsActive,
			},
			Statistics: PlayerStatisticsResponse{
				TotalMatches:       int(stats.TotalMatches),
				TotalMinutesPlayed: int(stats.TotalMinutesPlayed),
				TotalGoals:         int(stats.TotalGoals),
				TotalAssists:       int(stats.TotalAssists),
				TotalYellowCards:   int(stats.TotalYellowCards),
				TotalRedCards:      int(stats.TotalRedCards),
			},
			RecentMatches: recentMatches,
		},
	})
}
