package football_match

import (
	"context"
	"management-be/internal/model"
)

// GetMatchesByStatus gets football matches by status
func (i *impl) GetMatchesByStatus(ctx context.Context, status string) ([]model.FootballMatch, error) {
	return i.repo.FootballMatch().GetMatchesByStatus(ctx, status)
}
