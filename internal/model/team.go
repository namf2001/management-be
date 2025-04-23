package model

import (
	"time"
)

// Team represents a team entity
type Team struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	CompanyName   string     `json:"company_name"`
	ContactPerson string     `json:"contact_person"`
	ContactPhone  string     `json:"contact_phone"`
	ContactEmail  string     `json:"contact_email"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

// MatchHistory represents a team's match history
type MatchHistory struct {
	TotalMatches int     `json:"total_matches"`
	Wins         int     `json:"wins"`
	Losses       int     `json:"losses"`
	Draws        int     `json:"draws"`
	Matches      []Match `json:"match,omitempty"`
}
