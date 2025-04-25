package match

import (
	"context"
	"management-be/internal/model"
	"time"
)

// ListMatches returns all matches with optional filters
func (i impl) ListMatches(ctx context.Context, status string, startDate, endDate time.Time, opponentTeamID int) ([]model.Match, error) {
	// Call the repository method
	matches, err := i.repo.Match().ListMatches(ctx, status, startDate, endDate, opponentTeamID)
	if err != nil {
		return nil, err
	}

	return matches, nil
}
