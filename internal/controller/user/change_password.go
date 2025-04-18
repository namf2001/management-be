package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// UpdatePassword changes a user's password
func (i impl) UpdatePassword(ctx context.Context, userID int32, currentPassword, newPassword string) error {
	// Get the user
	userRepo := i.repo.User()
	user, err := userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrDatabase) {
			return ErrUserNotFound
		}
		return err
	}

	// Verify current password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
	if err != nil {
		return ErrInvalidCredentials
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update the password
	err = userRepo.UpdatePassword(ctx, userID, string(hashedPassword))
	if err != nil {
		return err
	}
	// For now, we'll just return a placeholder error
	return errors.New("update password not implemented")
}
