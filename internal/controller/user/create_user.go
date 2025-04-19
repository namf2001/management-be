package user

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user account
func (i impl) CreateUser(ctx context.Context, username, email, password string) (model.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(ErrHashingPassword)
	}

	// Create the user
	userRepo := i.repo.User()
	user, err := userRepo.CreateUser(ctx, username, email, string(hashedPassword))
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
