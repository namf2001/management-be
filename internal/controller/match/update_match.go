package match

import (
	"context"
	"errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// ErrMatchNotFound is returned when a match is not found
var ErrMatchNotFound = errors.New("match not found")

// ErrMatchNotUpdated is returned when a match cannot be updated
var ErrMatchNotUpdated = errors.New("match could not be updated")

// UpdateMatch updates an existing match with transaction support
func (i impl) UpdateMatch(ctx context.Context, id int, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, ourScore, opponentScore int, status, notes string) (model.Match, error) {
	var match model.Match

	// Execute the update operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		matchRepo := i.repo.Match()

		// Check if match exists
		_, err := matchRepo.GetMatch(ctx, id)
		if err != nil {
			return ErrMatchNotFound
		}

		// Update match
		updatedMatch, err := matchRepo.UpdateMatch(ctx, id, opponentTeamID, matchDate, venue, isHomeGame, ourScore, opponentScore, status, notes)
		if err != nil {
			return ErrMatchNotUpdated
		}

		match = updatedMatch
		return nil
	})

	if err != nil {
		return model.Match{}, err
	}

	return match, nil
}
