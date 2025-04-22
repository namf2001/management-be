package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetTeamByID retrieves a team by its ID
func (i impl) GetTeamByID(ctx context.Context, id int) (model.Team, error) {
	// Validate input
	if id <= 0 {
		return model.Team{}, pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Get team from repository
	teamRepo := i.repo.Team()
	team, err := teamRepo.GetTeamByID(ctx, id)
	if err != nil {
		return model.Team{}, pkgerrors.WithStack(ErrTeamNotFoundByID)
	}

	return team, nil
}
