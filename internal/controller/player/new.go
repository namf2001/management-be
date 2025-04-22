package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
	"time"
)

// Controller defines the interface for player controller operations
type Controller interface {
	CreatePlayer(ctx context.Context, departmentID int, fullName, position string, jerseyNumber int32, dateOfBirth *time.Time, heightCm, weightKg int32, phone, email string, isActive bool) (model.Player, error)
	GetPlayerByID(ctx context.Context, id int) (model.Player, error)
	GetAllPlayers(ctx context.Context, page, limit int) ([]model.Player, int, error)
	UpdatePlayer(ctx context.Context, id, departmentID int, fullName, position string, jerseyNumber int32, dateOfBirth *time.Time, heightCm, weightKg int32, phone, email string, isActive bool) (model.Player, error)
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
