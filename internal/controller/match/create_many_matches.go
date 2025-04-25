package match

import (
	"context"
	"errors"
	"management-be/internal/model"
)

// ErrMatchesNotCreated is returned when matches cannot be created
var ErrMatchesNotCreated = errors.New("matches could not be created")

// CreateManyMatches creates multiple matches at once
func (i impl) CreateManyMatches(ctx context.Context, matches []model.Match) ([]model.Match, error) {
	// Call the repository method
	createdMatches, err := i.repo.Match().CreateManyMatches(ctx, matches)
	if err != nil {
		return nil, ErrMatchesNotCreated
	}

	return createdMatches, nil
}
