package player

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/playerstatistic"
)

// GetPlayerStatistics retrieves the statistics of a player by their ID from the database.
func (i impl) GetPlayerStatistics(ctx context.Context, id int) (model.PlayerStatistic, error) {
	// Get player statistics using ent client
	stat, err := i.entClient.PlayerStatistic.Query().
		Where(playerstatistic.PlayerID(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.PlayerStatistic{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.PlayerStatistic{}, pkgerrors.WithStack(ErrDatabase)
	}

	// Convert ent.PlayerStatistic to model.PlayerStatistic
	return model.PlayerStatistic{
		ID:                 stat.ID,
		PlayerID:           stat.PlayerID,
		TotalMatches:       stat.TotalMatches,
		TotalMinutesPlayed: stat.TotalMinutesPlayed,
		TotalGoals:         stat.TotalGoals,
		TotalAssists:       stat.TotalAssists,
		TotalYellowCards:   stat.TotalYellowCards,
		TotalRedCards:      stat.TotalRedCards,
		CreatedAt:          stat.CreatedAt,
		UpdatedAt:          stat.UpdatedAt,
	}, nil
}
