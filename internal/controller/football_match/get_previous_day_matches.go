package football_match

import (
	"context"
	"management-be/internal/model"
)

// GetPreviousDayMatches gets football matches from the previous day
func (i *impl) GetPreviousDayMatches(ctx context.Context) ([]model.FootballMatch, error) {
	return i.repo.FootballMatch().GetPreviousDayMatches(ctx)
}
