package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchesgateway"
	"time"
)

// GetMatchesByDateRange gets football matches within a date range
func (i impl) GetMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatch, error) {
	matches, err := i.entClient.MatchesGateway.
		Query().
		Where(
			matchesgateway.And(
				matchesgateway.MatchDateGTE(startDate),
				matchesgateway.MatchDateLTE(endDate),
			),
		).
		Order(matchesgateway.ByMatchDate()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntsToModels(matches), nil
}
