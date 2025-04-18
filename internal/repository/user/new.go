package user

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

type Repository interface {
	CreateUser(ctx context.Context, username, email, password string) (model.User, error)
	GetUserByID(ctx context.Context, id int32) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type impl struct {
	entClient *ent.Client
}

func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}
