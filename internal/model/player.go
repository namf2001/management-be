package model

import (
	"time"
)

// Player represents a player entity
type Player struct {
	ID           int        `json:"id"`
	DepartmentID int        `json:"department_id"`
	FullName     string     `json:"full_name"`
	JerseyNumber *int32     `json:"jersey_number,omitempty"`
	Position     string     `json:"position"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty"`
	HeightCM     *int32     `json:"height_cm,omitempty"`
	WeightKG     *int32     `json:"weight_kg,omitempty"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
