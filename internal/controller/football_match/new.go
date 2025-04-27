package football_match

import (
	"context"
	"management-be/internal/gateway/football_matchs"
	"management-be/internal/model"
	"management-be/internal/repository"
	"time"
)

// Controller defines the interface for football match controller operations
type Controller interface {
	// FetchAndSavePreviousDayMatches fetches matches from the previous day and saves them to the database
	FetchAndSavePreviousDayMatches(ctx context.Context) error

	// GetMatchesByCompetition gets football matches by competition name
	GetMatchesByCompetition(ctx context.Context, competitionName string) ([]model.FootballMatch, error)

	// GetMatchesByDateRange gets football matches within a date range
	GetMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatch, error)

	// GetMatchesByStatus gets football matches by status
	GetMatchesByStatus(ctx context.Context, status string) ([]model.FootballMatch, error)

	// GetPreviousDayMatches gets football matches from the previous day
	GetPreviousDayMatches(ctx context.Context) ([]model.FootballMatch, error)
}

type impl struct {
	repo    repository.Registry
	gateway football_matchs.Gateway
}

// NewController creates a new football match controller
func NewController(repo repository.Registry, gateway football_matchs.Gateway) Controller {
	return &impl{
		repo:    repo,
		gateway: gateway,
	}
}
