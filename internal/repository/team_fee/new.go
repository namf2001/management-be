package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// Repository defines the interface for team fee repository operations
type Repository interface {
	// GetTeamFeeByID retrieves a team fee by its ID
	GetTeamFeeByID(ctx context.Context, id int) (model.TeamFee, error)
	// ListTeamFees returns all team fees with optional filters
	ListTeamFees(ctx context.Context, startDate, endDate *time.Time) ([]model.TeamFee, error)
	// CreateTeamFee creates a new team fee
	CreateTeamFee(ctx context.Context, teamFee CreateTeamFeeInput) (model.TeamFee, error)
	// UpdateTeamFee updates an existing team fee
	UpdateTeamFee(ctx context.Context, id int, teamFee UpdateTeamFeeInput) (model.TeamFee, error)
	// DeleteTeamFee deletes a team fee by ID
	DeleteTeamFee(ctx context.Context, id int) error
}

type impl struct {
	entClient *ent.Client
}

// NewRepository creates a new team fee repository
func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}
