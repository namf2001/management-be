package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchesgateway"
)

// GetMatchesByCompetition gets football matches by competition name
func (i impl) GetMatchesByCompetition(ctx context.Context, competitionName string) ([]model.FootballMatch, error) {
	matches, err := i.entClient.MatchesGateway.
		Query().
		Where(matchesgateway.CompetitionNameEQ(competitionName)).
		Order(matchesgateway.ByMatchDate()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntsToModels(matches), nil
}
