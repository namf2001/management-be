package team_fee

import "errors"

var (
	// ErrNotFound is returned when a team fee is not found
	ErrNotFound = errors.New("team fee not found")
)
