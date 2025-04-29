package player

import (
	"context"
	"management-be/internal/model"
)

// GetAllPlayers retrieves all players from the database with pagination and filters.
func (i impl) GetAllPlayers(ctx context.Context, page, limit int, departmentID *int, isActive *bool, position string) ([]model.Player, int, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Get players from repository with filters
	players, total, err := i.repo.Player().GetAllPlayers(ctx, offset, limit, departmentID, isActive, position)
	if err != nil {
		return nil, 0, err
	}

	return players, total, nil
}
