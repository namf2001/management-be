package user

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"management-be/internal/repository/ent"
)

// UpdatePassword changes a user's password
func (i impl) UpdatePassword(ctx context.Context, userID int, currentPassword, newPassword string) error {
	// Get the user
	user, err := i.repo.User().GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrDatabase) {
			return err
		}
		return pkgerrors.WithStack(ErrUserNotFoundByID)
	}

	// Verify the current password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
	if err != nil {
		return pkgerrors.WithStack(ErrInvalidCredentials)
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return pkgerrors.WithStack(ErrHashingPassword)
	}

	// Execute the update operation within a transaction
	err = i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		// Update the password
		err := i.repo.User().UpdatePassword(ctx, userID, string(hashedPassword))
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
