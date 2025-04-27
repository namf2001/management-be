package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchesgateway"
	"time"
)

// GetPreviousDayMatches gets football matches from the previous day
func (i impl) GetPreviousDayMatches(ctx context.Context) ([]model.FootballMatch, error) {
	// Calculate yesterday's date
	yesterday := time.Now().AddDate(0, 0, -1)
	// Set to beginning of day
	startOfDay := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
	// Set to end of day
	endOfDay := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 999999999, yesterday.Location())

	matches, err := i.entClient.MatchesGateway.
		Query().
		Where(
			matchesgateway.And(
				matchesgateway.MatchDateGTE(startOfDay),
				matchesgateway.MatchDateLTE(endOfDay),
			),
		).
		Order(matchesgateway.ByMatchDate()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntsToModels(matches), nil
}
