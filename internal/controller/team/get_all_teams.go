package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetAllTeams retrieves all teams with pagination
func (i impl) GetAllTeams(ctx context.Context, page, limit int) ([]model.Team, int, error) {
	// Validate input
	if page <= 0 || limit <= 0 {
		return nil, 0, pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Get teams from repository
	teamRepo := i.repo.Team()
	teams, total, err := teamRepo.GetAllTeams(ctx, page, limit)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(ErrDatabase)
	}

	return teams, total, nil
}
