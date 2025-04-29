package match

import "errors"

var (
	// ErrMatchNotFound is returned when a match is not found in the database.
	ErrMatchNotFound = errors.New("match not found")
	// ErrMatchAlreadyExists is returned when a match already exists in the database.
	ErrMatchAlreadyExists = errors.New("match already exists")
	// ErrDatabase is returned when there is a database error.
	ErrDatabase = errors.New("database error")
	// ErrInvalidMatchID is returned when the match ID is invalid.
	ErrInvalidMatchID = errors.New("invalid match ID")
	// ErrNotFound is returned when a match is not found.
	ErrNotFound = errors.New("not found")
)
