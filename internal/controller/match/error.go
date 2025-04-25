package match

import "errors"

var (
	// ErrMatchStatisticsNotFound is returned when match statistics cannot be found
	ErrMatchStatisticsNotFound = errors.New("match statistics not found")
	// ErrMatchNotFound is returned when a match is not found
	ErrMatchNotFound = errors.New("match not found")
	// ErrMatchNotUpdated is returned when a match cannot be updated
	ErrMatchNotUpdated = errors.New("match could not be updated")
	// ErrInvalidPlayerData is returned when player data is invalid
	ErrInvalidPlayerData = errors.New("invalid player data")
)
