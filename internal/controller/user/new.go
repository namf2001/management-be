package user

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
)

type Controller interface {
	CreateUser(ctx context.Context, username, email, password string) (int, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)
}

type impl struct {
	repo repository.Registry
}

func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
