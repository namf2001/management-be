package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/playerstatistic"
)

// GetPlayerStatistics retrieves the statistics of a player by their ID from the database.
func (i impl) GetPlayerStatistics(ctx context.Context, id int) (model.PlayerStatistic, error) {
	// Get player statistics using ent client
	stat, err := i.entClient.PlayerStatistic.Query().
		Where(playerstatistic.PlayerID(id)).
		Only(ctx)
	if err != nil {
		return model.PlayerStatistic{}, err
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
