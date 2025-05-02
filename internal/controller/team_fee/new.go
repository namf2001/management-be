package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
	"time"
)

// Controller defines the interface for team fee controller operations
type Controller interface {
	// ListTeamFees returns all team fees with optional filters
	ListTeamFees(ctx context.Context, startDate, endDate *time.Time) ([]model.TeamFee, model.TeamFeeSummary, error)
	// GetTeamFee returns a team fee by ID
	GetTeamFee(ctx context.Context, id int) (model.TeamFee, error)
	// CreateTeamFee creates a new team fee
	CreateTeamFee(ctx context.Context, input CreateTeamFeeInput) (model.TeamFee, error)
	// UpdateTeamFee updates an existing team fee
	UpdateTeamFee(ctx context.Context, id int, input UpdateTeamFeeInput) (model.TeamFee, error)
	// DeleteTeamFee deletes a team fee by ID
	DeleteTeamFee(ctx context.Context, id int) error
	// GetTeamFeeStatistics returns statistics about team fees
	GetTeamFeeStatistics(ctx context.Context, year *int) (model.TeamFeeStatistics, error)
}

type impl struct {
	repo repository.Registry
}

// NewController creates a new team fee controller
func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
