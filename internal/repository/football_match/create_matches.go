package football_match

import (
	"context"
	"management-be/internal/model"
)

// CreateMatches saves multiple football matches to the database
func (i impl) CreateMatches(ctx context.Context, matches []model.FootballMatch) ([]model.FootballMatch, error) {
	var createdMatches []model.FootballMatch

	for _, match := range matches {
		createdMatch, err := i.CreateMatch(ctx, match)
		if err != nil {
			return nil, err
		}
		createdMatches = append(createdMatches, createdMatch)
	}

	return createdMatches, nil
}
