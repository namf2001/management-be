package match

import (
	"context"
	"management-be/internal/model"
)

// GetMatch returns a match by ID with detailed information
func (i impl) GetMatch(ctx context.Context, id int) (model.Match, error) {
	// Call the repository method
	match, err := i.repo.Match().GetMatch(ctx, id)
	if err != nil {
		return model.Match{}, err
	}

	return match, nil
}
