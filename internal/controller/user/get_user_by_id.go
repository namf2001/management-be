package user

import (
	"context"
	"management-be/internal/model"
)

// GetUserByID retrieves a user by their ID.
func (i impl) GetUserByID(ctx context.Context, id int) (model.User, error) {
	user, err := i.repo.User().GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
