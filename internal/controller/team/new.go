package team

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
)

// Controller defines the interface for team controller operations
type Controller interface {
	CreateTeam(ctx context.Context, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error)
	GetTeamByID(ctx context.Context, id int) (model.Team, error)
	GetAllTeams(ctx context.Context, page, limit int) ([]model.Team, int, error)
	UpdateTeam(ctx context.Context, id int, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error)
	DeleteTeam(ctx context.Context, id int) error
	GetTeamStatistics(ctx context.Context, id int) (model.MatchHistory, error)
}

type impl struct {
	repo repository.Registry
}

// NewController creates a new team controller
func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
