package user

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserNotFound       = errors.New("user not found")
)

// Login authenticates a user and returns a JWT token
func (i impl) Login(ctx context.Context, username, password string) (string, model.User, error) {
	user, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		if !errors.Is(err, ErrUserNotFoundByUsername) {
			return "", model.User{}, pkgerrors.WithStack(ErrUserNotFound)
		}
		return "", model.User{}, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", model.User{}, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", model.User{}, err
	}

	return token, user, nil
}

// generateJWT creates a new JWT token for the given user ID
func generateJWT(userID int) (string, error) {
	// In a real implementation, you would get the secret from environment variables
	secret := []byte("your-secret-key")

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiration
		"iat":     time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
