package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// Repository defines the interface for team repository operations
type Repository interface {
	CreatePlayer(ctx context.Context, input InputPlayer) (model.Player, error)
	GetPlayerByID(ctx context.Context, id int) (model.Player, error)
	GetAllPlayers(ctx context.Context, offset, limit int, departmentID *int, isActive *bool, position string) ([]model.Player, int, error)
	UpdatePlayer(ctx context.Context, id int, input InputPlayer) (model.Player, error)
	DeletePlayer(ctx context.Context, id int) error
	GetPlayerStatistics(ctx context.Context, id int) (model.PlayerStatistic, error)
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
