package user

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent/user"
)

// GetUserByEmail retrieves a user by their email address.
func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	foundUser, err := i.entClient.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(ErrDatabase)
	}

	return model.User{
		ID:       foundUser.ID,
		Username: foundUser.Username,
		Email:    foundUser.Email,
	}, nil
}
