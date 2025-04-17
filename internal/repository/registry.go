package repository

import (
	"management-be/internal/repository/ent"
	"management-be/internal/repository/user"
)

type Registry interface {
	User() user.Repository
}

type impl struct {
	entConn *ent.Client
	user    user.Repository
}

func NewRegistry(entConn *ent.Client) Registry {
	return &impl{
		entConn: entConn,
		user:    user.NewRepository(entConn),
	}
}

func (i *impl) User() user.Repository {
	return i.user
}
