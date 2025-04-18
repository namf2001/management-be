package user

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/user"
)

// GetUserByEmail retrieves a user by their email address.
func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	only, err := i.entClient.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       int64(only.ID),
		Username: only.Username,
		Email:    only.Email,
	}, nil
}
