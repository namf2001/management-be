package v1

import "errors"

var (
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserNotValid           = errors.New("user not valid")
	ErrUserNotCreated         = errors.New("user not created")
	ErrUserNotUpdated         = errors.New("user not updated")
	ErrUserNotDeleted         = errors.New("user not deleted")
	ErrUserNotActivated       = errors.New("user not activated")
	ErrUserNotDeactivated     = errors.New("user not deactivated")
	ErrUserNotLoggedIn        = errors.New("user not logged in")
	ErrUserNotLoggedOut       = errors.New("user not logged out")
	ErrUserNotFoundByEmail    = errors.New("user not found by email")
	ErrUserNotFoundByUsername = errors.New("user not found by username")
	ErrUserNotFoundByID       = errors.New("user not found by id")
	ErrDatabase               = errors.New("database error")
)

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
