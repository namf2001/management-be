package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchesgateway"
)

// GetMatchesByStatus gets football matches by status
func (i impl) GetMatchesByStatus(ctx context.Context, status string) ([]model.FootballMatch, error) {
	matches, err := i.entClient.MatchesGateway.
		Query().
		Where(matchesgateway.StatusEQ(status)).
		Order(matchesgateway.ByMatchDate()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntsToModels(matches), nil
}
