package model

import (
	"time"
)

type MatchStatus string

const (
	MatchStatusScheduled MatchStatus = "scheduled"
	MatchStatusCompleted MatchStatus = "completed"
	MatchStatusCancelled MatchStatus = "cancelled"
)

type Match struct {
	ID             int        `json:"id"`
	OpponentTeamID int        `json:"opponent_team_id"`
	MatchDate      time.Time  `json:"match_date"`
	Venue          string     `json:"venue"`
	IsHomeGame     bool       `json:"is_home_game"`
	OurScore       int32      `json:"our_score,omitempty"`
	OpponentScore  int32      `json:"opponent_score,omitempty"`
	Status         string     `json:"status"`
	Notes          string     `json:"notes"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

type MatchStatistics struct {
	MatchSummary struct {
		TotalPlayers       int32 `json:"total_players"`
		TotalMinutesPlayed int32 `json:"total_minutes_played"`
		TotalGoals         int32 `json:"total_goals"`
		TotalAssists       int32 `json:"total_assists"`
		TotalYellowCards   int32 `json:"total_yellow_cards"`
		TotalRedCards      int32 `json:"total_red_cards"`
	} `json:"match_summary"`
	PlayerPerformance []struct {
		PlayerID      int    `json:"player_id"`
		PlayerName    string `json:"player_name"`
		Position      string `json:"position"`
		MinutesPlayed int32  `json:"minutes_played"`
		GoalsScored   int32  `json:"goals_scored"`
		Assists       int32  `json:"assists"`
		YellowCards   int32  `json:"yellow_cards"`
		RedCard       bool   `json:"red_card"`
	} `json:"player_performance"`
}
