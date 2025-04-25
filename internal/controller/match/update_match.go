package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// UpdateMatch updates an existing match with transaction support
func (i impl) UpdateMatch(ctx context.Context, id int, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, ourScore, opponentScore int32, status, notes string) (model.Match, error) {
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
