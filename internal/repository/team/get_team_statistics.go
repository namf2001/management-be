package team

import (
	"context"
	"entgo.io/ent/dialect/sql"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetTeamStatistics retrieves statistics for a team
func (i impl) GetTeamStatistics(ctx context.Context, id int) (model.MatchHistory, error) {
	// Check if team exists
	exists, err := i.entClient.Team.Query().
		Where(sql.FieldEQ("id", id)).
		Exist(ctx)

	if err != nil {
		return model.MatchHistory{}, pkgerrors.WithStack(ErrDatabase)
	}

	if !exists {
		return model.MatchHistory{}, pkgerrors.WithStack(ErrNotFound)
	}

	// For now, return empty statistics since we don't have match data yet
	// In a real implementation, this would query the matches table and calculate statistics
	return model.MatchHistory{
		TotalMatches: 0,
		Wins:         0,
		Losses:       0,
		Draws:        0,
		Matches:      []model.Match{},
	}, nil
}