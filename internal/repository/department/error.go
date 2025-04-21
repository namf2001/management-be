package department

import (
	"github.com/pkg/errors"
)

var (
	// ErrDatabase is returned when a database error occurs.
	ErrDatabase = errors.New("database error")

	// ErrNotFound is returned when a department is not found.
	ErrNotFound = errors.New("department not found")
)