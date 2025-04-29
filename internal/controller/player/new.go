package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
)

// Controller defines the interface for player controller operations
type Controller interface {
	CreatePlayer(ctx context.Context, input InputPlayerController) (model.Player, error)
	GetPlayerByID(ctx context.Context, id int) (model.Player, error)
	GetAllPlayers(ctx context.Context, page, limit int, departmentID *int, isActive *bool, position string) ([]model.Player, int, error)
	UpdatePlayer(ctx context.Context, id int, input InputPlayerController) (model.Player, error)
	DeletePlayer(ctx context.Context, id int) error
	GetPlayerStatistics(ctx context.Context, id int) (model.PlayerStatistic, error)
}

type impl struct {
	repo repository.Registry
}

// NewController creates a new player controller
func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
