package user

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

type Repository interface {
	CreateUser(ctx context.Context, username, email, password string) (int, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)
}

type impl struct {
	entClient *ent.Client
}

func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}
