package team

import (
	"errors"
)

// ErrDatabase is returned when a database error occurs.
var ErrDatabase = errors.New("database error")

// ErrNotFound is returned when a team is not found.
var ErrNotFound = errors.New("team not found")