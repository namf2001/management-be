package user

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// CreateUser creates a new user in the database.
func (i impl) CreateUser(ctx context.Context, username, email, password, fullName string) (model.User, error) {
	createdUser, err := i.entClient.User.Create().
		SetUsername(username).
		SetEmail(email).
		SetPassword(password).
		SetFullName(fullName).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	return model.User{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		FullName: createdUser.FullName,
	}, nil
}
