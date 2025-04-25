package player_statistics

import (
	"context"
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
	// GetMatchPlayers returns players in a match
	GetMatchPlayers(ctx context.Context, matchID int) ([]ent.Player, error)
}

type impl struct {
	entClient *ent.Client
}

func (i impl) GetMatchPlayers(ctx context.Context, matchID int) ([]ent.Player, error) {
	//TODO implement me
	panic("implement me")
}

// NewRepository creates a new team repository
func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}
