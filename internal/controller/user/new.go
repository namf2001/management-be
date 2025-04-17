package user

import (
	"management-be/internal/repository"
	"net/http"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
}

type impl struct {
	repo repository.Registry
}

func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
