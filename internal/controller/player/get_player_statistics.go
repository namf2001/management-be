package player

import (
	"context"
	"management-be/internal/model"
)

// GetPlayerStatistics retrieves the statistics of a player by their ID from the database.
func (i impl) GetPlayerStatistics(ctx context.Context, id int) (model.PlayerStatistic, error) {
	return i.repo.Player().GetPlayerStatistics(ctx, id)
}
