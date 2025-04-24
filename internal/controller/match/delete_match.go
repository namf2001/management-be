package match

import (
	"context"
	"errors"
	"management-be/internal/repository/ent"
)

// ErrMatchNotDeleted is returned when a match cannot be deleted
var ErrMatchNotDeleted = errors.New("match could not be deleted")

// DeleteMatch deletes a match by ID with transaction support
func (i impl) DeleteMatch(ctx context.Context, id int) error {
	// Execute the delete operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		matchRepo := i.repo.Match()

		// Check if match exists
		_, err := matchRepo.GetMatch(ctx, id)
		if err != nil {
			return ErrMatchNotFound
		}

		// Delete match
		err = matchRepo.DeleteMatchByID(ctx, id)
		if err != nil {
			return ErrMatchNotDeleted
		}

		return nil
	})

	return err
}
