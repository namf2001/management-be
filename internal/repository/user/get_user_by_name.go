package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/user"
)

// GetUserByUsername retrieves a user by their username.
func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	foundUser, err := i.entClient.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.User{}, pkgerrors.WithStack(ErrUserNotFoundByUsername)
		}
		return model.User{}, pkgerrors.WithStack(ErrDatabase)
	}
	return model.User{
		ID:       foundUser.ID,
		Username: foundUser.Username,
		Email:    foundUser.Email,
		Password: foundUser.Password,
		FullName: foundUser.FullName,
	}, nil
}
