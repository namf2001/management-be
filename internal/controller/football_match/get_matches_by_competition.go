package football_match

import (
	"context"
	"management-be/internal/model"
)

// GetMatchesByCompetition gets football matches by competition name
func (i *impl) GetMatchesByCompetition(ctx context.Context, competitionName string) ([]model.FootballMatch, error) {
	return i.repo.FootballMatch().GetMatchesByCompetition(ctx, competitionName)
}
