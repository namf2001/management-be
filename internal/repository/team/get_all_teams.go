package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetAllTeams retrieves all teams with pagination
func (i impl) GetAllTeams(ctx context.Context, page, limit int) ([]model.Team, int, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Get total count
	total, err := i.entClient.Team.Query().Count(ctx)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Get teams with pagination
	teams, err := i.entClient.Team.Query().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	// Convert to model.Team
	result := make([]model.Team, len(teams))
	for idx, team := range teams {
		result[idx] = model.Team{
			ID:            team.ID,
			Name:          team.Name,
			CompanyName:   team.CompanyName,
			ContactPerson: team.ContactPerson,
			ContactPhone:  team.ContactPhone,
			ContactEmail:  team.ContactEmail,
			CreatedAt:     team.CreatedAt,
			UpdatedAt:     team.UpdatedAt,
		}
	}

	return result, total, nil
}
