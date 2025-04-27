package player

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetAllPlayers retrieves all players from the database with pagination.
func (i impl) GetAllPlayers(ctx context.Context, page, limit int) ([]model.Player, int, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Query players with pagination
	players, err := i.entClient.Player.Query().
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Get total count
	total, err := i.entClient.Player.Query().Count(ctx)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Convert ent.Player to model.Player
	result := make([]model.Player, len(players))
	for i, player := range players {
		result[i] = model.Player{
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
		}
	}

	return result, total, nil
}
