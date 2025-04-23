package model

import "time"

// PlayerStatistic represents a player's statistics in a season
type PlayerStatistic struct {
	ID                 int        `json:"id"`
	PlayerID           int        `json:"player_id"`
	TotalMatches       int32      `json:"total_matches"`
	TotalMinutesPlayed int32      `json:"total_minutes_played"`
	TotalGoals         int32      `json:"total_goals"`
	TotalAssists       int32      `json:"total_assists"`
	TotalYellowCards   int32      `json:"total_yellow_cards"`
	TotalRedCards      int32      `json:"total_red_cards"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}
