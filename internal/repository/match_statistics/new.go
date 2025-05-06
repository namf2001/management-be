package match_statistics

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
	// GetMatchPlayers returns players in a match
	GetMatchPlayers(ctx context.Context, matchID int) ([]model.MatchPlayer, error)
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
