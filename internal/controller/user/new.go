package user

import (
	"context"
	"management-be/internal/repository"
)

type Controller interface {
	CreateUser(ctx context.Context, username, email, password string) error
}

type impl struct {
	repo repository.Registry
}

func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
