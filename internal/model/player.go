package model

import (
	"time"
)

// Player represents a player entity
type Player struct {
	ID           int        `json:"id"`
	DepartmentID int        `json:"department_id,omitempty"`
	FullName     string     `json:"full_name"`
	JerseyNumber int32      `json:"jersey_number,omitempty"`
	Position     string     `json:"position"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty"`
	HeightCm     int32      `json:"height_cm,omitempty"`
	WeightKg     int32      `json:"weight_kg,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	Email        string     `json:"email,omitempty"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// PlayerStatistic represents a player's statistics
type PlayerStatistic struct {
	ID                 int       `json:"id"`
	PlayerID           int       `json:"player_id"`
	TotalMatches       int32     `json:"total_matches"`
	TotalMinutesPlayed int32     `json:"total_minutes_played"`
	TotalGoals         int32     `json:"total_goals"`
	TotalAssists       int32     `json:"total_assists"`
	TotalYellowCards   int32     `json:"total_yellow_cards"`
	TotalRedCards      int32     `json:"total_red_cards"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
