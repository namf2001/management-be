package player

import (
	"context"
	"management-be/internal/model"
)

// GetAllPlayers retrieves all players from the database with pagination.
func (i impl) GetAllPlayers(ctx context.Context, page, limit int) ([]model.Player, int, error) {
	return i.repo.Player().GetAllPlayers(ctx, page, limit)
}
