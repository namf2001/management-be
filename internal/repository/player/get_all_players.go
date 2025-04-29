package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/player"

	pkgerrors "github.com/pkg/errors"
)

// GetAllPlayers retrieves all players from the database with pagination and filters.
func (i impl) GetAllPlayers(ctx context.Context, page, pageSize int, departmentID *int, isActive *bool, position string) ([]model.Player, int, error) {
	// Handle invalid page numbers
	if page < 1 {
		page = 1 // Default to page 1 for invalid values
	}

	// Calculate offset based on page and pageSize
	offset := (page - 1) * pageSize

	// Start building the query
	query := i.entClient.Player.Query()

	// Apply filters if provided
	if departmentID != nil {
		query = query.Where(player.DepartmentID(*departmentID))
	}
	if isActive != nil {
		query = query.Where(player.IsActive(*isActive))
	}
	if position != "" {
		query = query.Where(player.Position(position))
	}

	// Get total count with the same filters before applying pagination
	totalQuery := query.Clone()
	total, err := totalQuery.Count(ctx)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Apply pagination
	query = query.Offset(offset).Limit(pageSize)

	// Query players
	players, err := query.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return []model.Player{}, total, nil
		}
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
