package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// UpdateTeam updates an existing team
func (i impl) UpdateTeam(ctx context.Context, id int, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error) {
	// Validate input
	if id <= 0 || name == "" || companyName == "" || contactPerson == "" {
		return model.Team{}, pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Update team in repository
	teamRepo := i.repo.Team()
	team, err := teamRepo.UpdateTeam(ctx, id, name, companyName, contactPerson, contactPhone, contactEmail)
	if err != nil {
		return model.Team{}, pkgerrors.WithStack(ErrTeamNotUpdated)
	}

	return team, nil
}
