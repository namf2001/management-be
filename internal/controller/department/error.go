package department

import "errors"

var (
	ErrDepartmentAlreadyExists  = errors.New("department already exists")
	ErrDepartmentNotValid       = errors.New("department not valid")
	ErrDepartmentNotCreated     = errors.New("department not created")
	ErrDepartmentNotUpdated     = errors.New("department not updated")
	ErrDepartmentNotDeleted     = errors.New("department not deleted")
	ErrDepartmentNotFoundByID   = errors.New("department not found by id")
	ErrDepartmentNotFoundByName = errors.New("department not found by name")
	ErrDatabase                 = errors.New("database error")
	ErrDepartmentNotFound       = errors.New("department not found")
)
