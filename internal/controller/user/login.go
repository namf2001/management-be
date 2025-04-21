package user

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
		return "", model.User{}, pkgerrors.WithStack(ErrTokenGeneration)
	}

	return token, user, nil
}

// generateJWT creates a new JWT token for the given user ID
func generateJWT(userID int) (string, error) {
	// Get the secret from environment variables
	secret := []byte(os.Getenv("JWT_SECRET"))

	// Parse expiration time from environment variable
	expStr := os.Getenv("JWT_EXPIRATION")
	expDuration, err := parseExpiration(expStr)
	if err != nil {
		// Default to 24 hours if parsing fails
		expDuration = time.Hour * 24
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expDuration).Unix(),
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

// parseExpiration parses the expiration time string (e.g., "24h") into a time.Duration
func parseExpiration(expStr string) (time.Duration, error) {
	return time.ParseDuration(expStr)
}
