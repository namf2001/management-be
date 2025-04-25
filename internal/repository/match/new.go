package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// Repository defines the interface for team repository operations
type Repository interface {
	// ListMatches returns all matches with optional filters
	ListMatches(ctx context.Context, status string, startDate, endDate time.Time, opponentTeamID int) ([]model.Match, error)
	// GetMatch returns a match by ID with detailed information
	GetMatch(ctx context.Context, id int) (model.Match, error)
	// CreateMatch creates a new match
	CreateMatch(ctx context.Context, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, notes string) (model.Match, error)
	// CreateManyMatches creates multiple matches at once
	CreateManyMatches(ctx context.Context, matches []model.Match) ([]model.Match, error)
	// UpdateMatch updates an existing match
	UpdateMatch(ctx context.Context, id int, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, ourScore, opponentScore int32, status, notes string) (model.Match, error)
	// DeleteMatchByID deletes a match by ID
	DeleteMatchByID(ctx context.Context, id int) error
	// UpdateMatchPlayers updates player participation in a match
	UpdateMatchPlayers(ctx context.Context, matchID int, players []model.MatchPlayer) error
	// GetMatchStatistics returns match statistics and summary
	GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error)
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
