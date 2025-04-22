package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
)

// DeleteTeam deletes a team by its ID
func (i impl) DeleteTeam(ctx context.Context, id int) error {
	// Validate input
	if id <= 0 {
		return pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Delete team from repository
	teamRepo := i.repo.Team()
	err := teamRepo.DeleteTeam(ctx, id)
	if err != nil {
		return pkgerrors.WithStack(ErrTeamNotDeleted)
	}

	return nil
}
