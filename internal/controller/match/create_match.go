package match

import (
	"context"
	"errors"
	"management-be/internal/model"
	"time"
)

// ErrMatchNotCreated is returned when a match cannot be created
var ErrMatchNotCreated = errors.New("match could not be created")

// CreateMatch creates a new match
func (i impl) CreateMatch(ctx context.Context, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, notes string) (model.Match, error) {
	// Call the repository method
	match, err := i.repo.Match().CreateMatch(ctx, opponentTeamID, matchDate, venue, isHomeGame, notes)
	if err != nil {
		return model.Match{}, ErrMatchNotCreated
	}

	return match, nil
}
