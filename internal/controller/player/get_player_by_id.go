package player

import (
	"context"
	"management-be/internal/model"
)

// GetPlayerByID retrieves a player by their ID from the database.
func (i impl) GetPlayerByID(ctx context.Context, id int) (model.Player, error) {
	player, err := i.repo.Player().GetPlayerByID(ctx, id)
	if err != nil {
		return model.Player{}, err
	}

	// Get department name
	department, err := i.repo.Department().GetDepartmentByID(ctx, player.DepartmentID)
	if err != nil {
		return model.Player{}, err
	}

	// Add department name to the player model
	player.DepartmentName = department.Name

	return player, nil
}
