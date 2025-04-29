package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetUserByID retrieves a user by their ID.
func (i impl) GetUserByID(ctx context.Context, id int) (model.User, error) {
	user, err := i.entClient.User.Get(ctx, id)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}, nil
}
