package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/player"
)

// GetAllPlayers retrieves all players from the database with pagination and filters.
func (i impl) GetAllPlayers(ctx context.Context, page, limit int, departmentID *int, isActive *bool, position string) ([]model.Player, int, error) {
	// Get players from repository with filters
	filter := player.FilterGetAllPlayers{
		DepartmentID: departmentID,
		IsActive:     isActive,
		Position:     position,
		PageSize:     limit,
		Page:         page,
	}
	players, total, err := i.repo.Player().GetAllPlayers(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return players, total, nil
}
