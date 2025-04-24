package match

import (
	"context"
	"errors"
	"management-be/internal/model"
)

// ErrMatchStatisticsNotFound is returned when match statistics cannot be found
var ErrMatchStatisticsNotFound = errors.New("match statistics not found")

// GetMatchStatistics returns match statistics and summary
func (i impl) GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error) {
	// Check if match exists
	_, err := i.repo.Match().GetMatch(ctx, matchID)
	if err != nil {
		return model.MatchStatistics{}, ErrMatchNotFound
	}

	// Get match statistics
	stats, err := i.repo.Match().GetMatchStatistics(ctx, matchID)
	if err != nil {
		return model.MatchStatistics{}, ErrMatchStatisticsNotFound
	}

	return stats, nil
}
