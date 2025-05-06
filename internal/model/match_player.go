package model

import "time"

// MatchPlayer represents a player's performance in a match
type MatchPlayer struct {
	ID            int        `json:"id"`
	MatchID       int        `json:"match_id"`
	PlayerID      int        `json:"player_id"`
	MinutesPlayed int32      `json:"minutes_played"`
	GoalsScored   int32      `json:"goals_scored"`
	Assists       int32      `json:"assists"`
	YellowCards   int32      `json:"yellow_cards"`
	RedCard       bool       `json:"red_card"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
