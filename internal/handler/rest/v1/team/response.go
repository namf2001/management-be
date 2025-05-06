package team

// TeamResponse represents the response body for team operations
// @name TeamResponse
type TeamResponse struct {
	ID            int    `json:"id" example:"1"`
	Name          string `json:"name" example:"FC Barcelona"`
	CompanyName   string `json:"company_name" example:"FC Barcelona Sports Club"`
	ContactPerson string `json:"contact_person" example:"Joan Laporta"`
	ContactPhone  string `json:"contact_phone" example:"123456789"`
	ContactEmail  string `json:"contact_email" example:"contact@fcbarcelona.com"`
}

// MatchResponse represents a match in the team's history
// @name MatchResponse
type MatchResponse struct {
	MatchID       int    `json:"match_id" example:"1"`
	MatchDate     string `json:"match_date" example:"2024-06-01T15:00:00Z"`
	Venue         string `json:"venue" example:"Camp Nou"`
	OurScore      int32  `json:"our_score" example:"3"`
	OpponentScore int32  `json:"opponent_score" example:"1"`
	Status        string `json:"status" example:"completed"`
}

// MatchHistoryResponse represents the match history summary
// @name MatchHistoryResponse
type MatchHistoryResponse struct {
	TotalMatches int             `json:"total_matches" example:"10"`
	Wins         int             `json:"wins" example:"6"`
	Losses       int             `json:"losses" example:"3"`
	Draws        int             `json:"draws" example:"1"`
	Matches      []MatchResponse `json:"matches"`
}

// TeamWithStatsResponse represents the response body for team with statistics
// @name TeamWithStatsResponse
type TeamWithStatsResponse struct {
	ID            int                  `json:"id" example:"1"`
	Name          string               `json:"name" example:"FC Barcelona"`
	CompanyName   string               `json:"company_name" example:"FC Barcelona Sports Club"`
	ContactPerson string               `json:"contact_person" example:"Joan Laporta"`
	ContactPhone  string               `json:"contact_phone" example:"123456789"`
	ContactEmail  string               `json:"contact_email" example:"contact@fcbarcelona.com"`
	MatchHistory  MatchHistoryResponse `json:"match_history"`
}
