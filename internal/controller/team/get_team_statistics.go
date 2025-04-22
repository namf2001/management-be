package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetTeamStatistics retrieves statistics for a team
func (i impl) GetTeamStatistics(ctx context.Context, id int) (model.MatchHistory, error) {
	// Validate input
	if id <= 0 {
		return model.MatchHistory{}, pkgerrors.WithStack(ErrTeamNotValid)
	}

	// Get team statistics from repository
	teamRepo := i.repo.Team()
	stats, err := teamRepo.GetTeamStatistics(ctx, id)
	if err != nil {
		return model.MatchHistory{}, pkgerrors.WithStack(ErrTeamNotFoundByID)
	}

	return stats, nil
}
