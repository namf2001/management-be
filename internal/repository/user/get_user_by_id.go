package user

import (
	"context"
	"management-be/internal/model"
)

// GetUserByID retrieves a user by their ID.
func (i impl) GetUserByID(ctx context.Context, id int) (model.User, error) {
	only, err := i.entClient.User.Get(ctx, uint(id))
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       int64(only.ID),
		Username: only.Username,
		Email:    only.Email,
	}, nil
}
