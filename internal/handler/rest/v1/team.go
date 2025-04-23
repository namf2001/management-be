package v1

import (
	"management-be/internal/controller/team"
)

// TeamResponse represents the response format for team data
type TeamResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CompanyName   string `json:"company_name"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	ContactEmail  string `json:"contact_email"`
}

// TeamWithStatsResponse represents the response format for team data with match history
type TeamWithStatsResponse struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	CompanyName   string               `json:"company_name"`
	ContactPerson string               `json:"contact_person"`
	ContactPhone  string               `json:"contact_phone"`
	ContactEmail  string               `json:"contact_email"`
	MatchHistory  MatchHistoryResponse `json:"match_history"`
}

// MatchHistoryResponse represents the response format for match history
type MatchHistoryResponse struct {
	TotalMatches int             `json:"total_matches"`
	Wins         int             `json:"wins"`
	Losses       int             `json:"losses"`
	Draws        int             `json:"draws"`
	Matches      []MatchResponse `json:"matches,omitempty"`
}

// MatchResponse represents the response format for a match
type MatchResponse struct {
	MatchID       int    `json:"match_id"`
	MatchDate     string `json:"match_date"`
	Venue         string `json:"venue"`
	OurScore      int32  `json:"our_score"`
	OpponentScore int32  `json:"opponent_score"`
	Status        string `json:"status"`
}

// UpdateHandler to include team controller
func (h *Handler) UpdateTeamController(teamCtrl team.Controller) {
	h.teamCtrl = teamCtrl
}
