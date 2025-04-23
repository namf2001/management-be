package player

import (
	"context"
	"management-be/internal/model"
)

// GetPlayerByID retrieves a player by their ID from the database.
func (i impl) GetPlayerByID(ctx context.Context, id int) (model.Player, error) {
	player, err := i.entClient.Player.Get(ctx, id)
	if err != nil {
		return model.Player{}, err
	}

	return model.Player{
		ID:           player.ID,
		DepartmentID: player.DepartmentID,
		FullName:     player.FullName,
		Position:     player.Position,
		JerseyNumber: &player.JerseyNumber,
		DateOfBirth:  &player.DateOfBirth,
		HeightCM:     &player.HeightCm,
		WeightKG:     &player.WeightKg,
		Phone:        player.Phone,
		Email:        player.Email,
		IsActive:     player.IsActive,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}, nil
}
