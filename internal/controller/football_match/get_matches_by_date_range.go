package football_match

import (
	"context"
	"management-be/internal/model"
	"time"
)

// GetMatchesByDateRange gets football matches within a date range
func (i *impl) GetMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatch, error) {
	return i.repo.FootballMatch().GetMatchesByDateRange(ctx, startDate, endDate)
}
