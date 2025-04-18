package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent/user"
)

func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	foundUser, err := i.entClient.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return model.User{}, pkgerrors.WithStack(ErrUserNotFoundByUsername)
		}
		return model.User{}, pkgerrors.WithStack(ErrDatabase)
	}
	return model.User{
		ID:       foundUser.ID,
		Username: foundUser.Username,
		Email:    foundUser.Email,
	}, nil
}
