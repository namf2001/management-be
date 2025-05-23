package player

import (
	"context"
	"management-be/internal/model"
)

// GetPlayerByID retrieves a player by their ID from the database.
func (i impl) GetPlayerByID(ctx context.Context, id int) (model.Player, error) {
	return i.repo.Player().GetPlayerByID(ctx, id)
}
