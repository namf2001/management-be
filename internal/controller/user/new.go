package user

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
)

type Controller interface {
	GetUserByID(ctx context.Context, id int) (model.User, error)
	Login(ctx context.Context, username, password string) (string, model.User, error)
	CreateUser(ctx context.Context, username, email, password string) (model.User, error)
	UpdatePassword(ctx context.Context, userID int, currentPassword, newPassword string) error
}

type impl struct {
	repo repository.Registry
}

func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
