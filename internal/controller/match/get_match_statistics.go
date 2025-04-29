package match

import (
	"context"
	"management-be/internal/model"
)

// GetMatchStatistics returns match statistics and summary
func (i impl) GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error) {
	// Check if match exists
	_, err := i.repo.Match().GetMatchByID(ctx, matchID)
	if err != nil {
		return model.MatchStatistics{}, err
	}

	// Get match statistics
	stats, err := i.repo.Match().GetMatchStatistics(ctx, matchID)
	if err != nil {
		return model.MatchStatistics{}, err
	}

	return stats, nil
}
