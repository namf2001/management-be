package department

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository"
)

type Controller interface {
	CreateDepartment(ctx context.Context, name, description string) (model.Department, error)
	GetDepartmentByID(ctx context.Context, id int) (model.Department, error)
	GetAllDepartments(ctx context.Context) ([]model.Department, error)
	UpdateDepartment(ctx context.Context, id int, name, description string) (model.Department, error)
	DeleteDepartment(ctx context.Context, id int) error
}

type impl struct {
	repo repository.Registry
}

func NewController(repo repository.Registry) Controller {
	return &impl{
		repo: repo,
	}
}
