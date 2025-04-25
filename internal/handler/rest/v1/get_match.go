package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// MatchPlayerResponse represents a player's performance in a match
type MatchPlayerResponse struct {
	PlayerID      int    `json:"player_id"`
	PlayerName    string `json:"player_name"`
	JerseyNumber  int    `json:"jersey_number"`
	Position      string `json:"position"`
	MinutesPlayed int    `json:"minutes_played"`
	GoalsScored   int    `json:"goals_scored"`
	Assists       int    `json:"assists"`
	YellowCards   int    `json:"yellow_cards"`
	RedCard       bool   `json:"red_card"`
}

// GetMatchResponse represents the response for getting a match by ID
type GetMatchResponse struct {
	ID            int                   `json:"id"`
	OpponentTeam  TeamDetailResponse    `json:"opponent_team"`
	MatchDate     time.Time             `json:"match_date"`
	Venue         string                `json:"venue"`
	IsHomeGame    bool                  `json:"is_home_game"`
	OurScore      int32                 `json:"our_score,omitempty"`
	OpponentScore int32                 `json:"opponent_score,omitempty"`
	Status        string                `json:"status"`
	Notes         string                `json:"notes"`
	Players       []MatchPlayerResponse `json:"players,omitempty"`
}

// TeamDetailResponse represents the detailed team information
type TeamDetailResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CompanyName   string `json:"company_name"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
}

// GetMatch handles the request to get a match by ID
func (h Handler) GetMatch(ctx *gin.Context) {
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
	match, err := h.matchCtrl.GetMatch(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Match not found",
		})
		return
	}

	// Prepare response
	// Note: In a real implementation, you would fetch the team details and match players
	response := GetMatchResponse{
		ID: match.ID,
		OpponentTeam: TeamDetailResponse{
			ID:            match.OpponentTeamID,
			Name:          "", // Need to fetch team details
			CompanyName:   "",
			ContactPerson: "",
			ContactPhone:  "",
		},
		MatchDate:     match.MatchDate,
		Venue:         match.Venue,
		IsHomeGame:    match.IsHomeGame,
		OurScore:      match.OurScore,
		OpponentScore: match.OpponentScore,
		Status:        match.Status,
		Notes:         match.Notes,
		Players:       []MatchPlayerResponse{}, // Need to fetch match players
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
