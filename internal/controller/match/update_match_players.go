package match

import (
	"context"
	"errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// ErrMatchPlayersNotUpdated is returned when match players cannot be updated
var ErrMatchPlayersNotUpdated = errors.New("match players could not be updated")

// UpdateMatchPlayers updates player participation in a match with transaction support
func (i impl) UpdateMatchPlayers(ctx context.Context, matchID int, players []model.MatchPlayer) error {
	// Execute the update operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		matchRepo := i.repo.Match()

		// Check if match exists
		_, err := matchRepo.GetMatchByID(ctx, matchID)
		if err != nil {
			return ErrMatchNotFound
		}

		// Update match players
		err = matchRepo.UpdateMatchPlayers(ctx, matchID, players)
		if err != nil {
			return ErrMatchPlayersNotUpdated
		}

		return nil
	})

	return err
}
