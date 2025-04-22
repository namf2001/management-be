package team

import "errors"

var (
	ErrTeamAlreadyExists  = errors.New("team already exists")
	ErrTeamNotValid       = errors.New("team not valid")
	ErrTeamNotCreated     = errors.New("team not created")
	ErrTeamNotUpdated     = errors.New("team not updated")
	ErrTeamNotDeleted     = errors.New("team not deleted")
	ErrTeamNotFoundByID   = errors.New("team not found by id")
	ErrTeamNotFoundByName = errors.New("team not found by name")
	ErrDatabase           = errors.New("database error")
	ErrTeamNotFound       = errors.New("team not found")
)
