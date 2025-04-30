package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
	"time"
)

// Controller defines the interface for match controller operations
type Controller interface {
	// ListMatches returns all matches with optional filters
	ListMatches(ctx context.Context, status string, startDate, endDate time.Time, opponentTeamID int) ([]model.Match, error)
	// GetMatch returns a match by ID with detailed information
	GetMatch(ctx context.Context, id int) (model.Match, error)
	// CreateMatch creates a new match
	CreateMatch(ctx context.Context, input CreateMatchInput) (model.Match, error)
	// CreateManyMatches creates multiple matches at once
	CreateManyMatches(ctx context.Context, matches []model.Match) ([]model.Match, error)
	// UpdateMatch updates an existing match
	UpdateMatch(ctx context.Context, id int, input UpdateMatchInput) (model.Match, error)
	// DeleteMatch deletes a match by ID
	DeleteMatch(ctx context.Context, id int) error
	// UpdateMatchPlayers updates player participation in a match
	UpdateMatchPlayers(ctx context.Context, matchID int, players []model.MatchPlayer) error
	// GetMatchStatistics returns match statistics and summary
	GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error)
}

type impl struct {
	repo repository.Registry
}

// NewController creates a new match controller
func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
