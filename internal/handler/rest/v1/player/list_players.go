package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListPlayersResponse represents the response for listing players
// @name ListPlayersResponse
type ListPlayersResponse struct {
	Players []PlayerListItem `json:"players"`
	Total   int              `json:"total" example:"50"`
}

// PlayerListItem represents a player in the list
// @name PlayerListItem
type PlayerListItem struct {
	ID             int              `json:"id" example:"1"`
	DepartmentID   int              `json:"department_id" example:"1"`
	DepartmentName string           `json:"department_name" example:"Engineering"`
	FullName       string           `json:"full_name" example:"John Doe"`
	JerseyNumber   int              `json:"jersey_number" example:"10"`
	Position       string           `json:"position" example:"Forward"`
	DateOfBirth    string           `json:"date_of_birth,omitempty" example:"1990-01-01"`
	HeightCm       int              `json:"height_cm,omitempty" example:"180"`
	WeightKg       int              `json:"weight_kg,omitempty" example:"75"`
	Phone          string           `json:"phone,omitempty" example:"+1234567890"`
	Email          string           `json:"email,omitempty" example:"john.doe@example.com"`
	IsActive       bool             `json:"is_active" example:"true"`
	Statistics     PlayerBasicStats `json:"statistics"`
}

// PlayerBasicStats represents basic statistics for a player
// @name PlayerBasicStats
type PlayerBasicStats struct {
	TotalMatches       int `json:"total_matches" example:"10"`
	TotalMinutesPlayed int `json:"total_minutes_played" example:"900"`
	TotalGoals         int `json:"total_goals" example:"5"`
	TotalAssists       int `json:"total_assists" example:"3"`
}

// ListPlayers
// @Summary      List all players
// @Description  Get a list of all players with their basic statistics
// @Tags         players
// @Accept       json
// @Produce      json
// @Success      200  {object}  ListPlayersResponse
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/players [get]
func (h Handler) ListPlayers(ctx *gin.Context) {
	// Get query parameters
	departmentIDStr := ctx.Query("department_id")
	isActiveStr := ctx.Query("is_active")
	position := ctx.Query("position")
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	// Parse pagination params
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse department ID if provided
	var departmentID *int
	if departmentIDStr != "" {
		id, err := strconv.Atoi(departmentIDStr)
		if err == nil {
			departmentID = &id
		}
	}

	// Parse isActive if provided
	var isActive *bool
	if isActiveStr != "" {
		active, err := strconv.ParseBool(isActiveStr)
		if err == nil {
			isActive = &active
		}
	}

	// Get all players from controller with filters
	players, total, err := h.playerCtrl.GetAllPlayers(ctx.Request.Context(), page, limit, departmentID, isActive, position)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch players",
		})
		return
	}

	// Convert players to response format
	response := make([]PlayerListItem, len(players))
	for i, player := range players {
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

		// Get basic statistics for each player
		stats, _ := h.playerCtrl.GetPlayerStatistics(ctx.Request.Context(), player.ID)

		// Set player data for response
		response[i] = PlayerListItem{
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
			Statistics: PlayerBasicStats{
				TotalMatches:       int(stats.TotalMatches),
				TotalMinutesPlayed: int(stats.TotalMinutesPlayed),
				TotalGoals:         int(stats.TotalGoals),
				TotalAssists:       int(stats.TotalAssists),
			},
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": ListPlayersResponse{
			Players: response,
			Total:   total,
		},
	})
}
