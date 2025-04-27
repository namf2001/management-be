package player

import (
	"github.com/pkg/errors"
)

var (
	// ErrDatabase is returned when a database error occurs.
	ErrDatabase = errors.New("database error")

	// ErrNotFound is returned when a player is not found.
	ErrNotFound = errors.New("player not found")
)
