package team

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
	CreateTeam(ctx context.Context, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error)
	GetTeamByID(ctx context.Context, id int) (model.Team, error)
	GetAllTeams(ctx context.Context, page, limit int) ([]model.Team, int, error)
	UpdateTeam(ctx context.Context, id int, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error)
	DeleteTeam(ctx context.Context, id int) error
	GetTeamStatistics(ctx context.Context, id int) (model.MatchHistory, error)
}

type impl struct {
	entClient *ent.Client
}

// NewRepository creates a new team repository
func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}