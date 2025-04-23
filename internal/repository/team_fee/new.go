package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
	// GetTeamFee retrieves a team fee by its ID
	GetTeamFee(ctx context.Context, id int) (model.TeamFee, error)
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
