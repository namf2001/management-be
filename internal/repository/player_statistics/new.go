package player_statistics

import (
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
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
