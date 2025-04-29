package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/player"

	pkgerrors "github.com/pkg/errors"
)

type FilterGetAllPlayers struct {
	DepartmentID *int
	IsActive     *bool
	Position     string
	PageSize     int
	Page         int
}

// GetAllPlayers retrieves all players from the database with pagination and filters.
func (i impl) GetAllPlayers(ctx context.Context, filter FilterGetAllPlayers) ([]model.Player, int, error) {
	// Handle invalid page numbers
	if filter.Page < 1 {
		filter.Page = 1 // Default to page 1 for invalid values
	}

	// Calculate offset based on page and pageSize
	offset := (filter.Page - 1) * filter.PageSize

	// Start building the query
	query := i.entClient.Player.Query()

	// Apply filters if provided
	if filter.DepartmentID != nil {
		query = query.Where(player.DepartmentID(*filter.DepartmentID))
	}
	if filter.IsActive != nil {
		query = query.Where(player.IsActive(*filter.IsActive))
	}
	if filter.Position != "" {
		query = query.Where(player.Position(filter.Position))
	}

	// Get total count with the same filters before applying pagination
	totalQuery := query.Clone()
	total, err := totalQuery.Count(ctx)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Apply pagination
	query = query.Offset(offset).Limit(filter.PageSize)

	// Query players
	players, err := query.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return []model.Player{}, total, pkgerrors.WithStack(ErrNotFound)
		}
		return nil, 0, pkgerrors.WithStack(err)
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
