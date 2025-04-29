package user

import "errors"

var (
	ErrNotFound         = errors.New("user not found")
	ErrUserNotFoundByID = errors.New("user not found by id")
	ErrDatabase         = errors.New("database error")
)
