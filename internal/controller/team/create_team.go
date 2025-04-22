package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// CreateTeam creates a new team
func (i impl) CreateTeam(ctx context.Context, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error) {
	// Validate input
	if name == "" || companyName == "" || contactPerson == "" {
		return model.Team{}, pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Create team in repository
	teamRepo := i.repo.Team()
	team, err := teamRepo.CreateTeam(ctx, name, companyName, contactPerson, contactPhone, contactEmail)
	if err != nil {
		return model.Team{}, pkgerrors.WithStack(ErrDatabase)
	}

	return team, nil
}
