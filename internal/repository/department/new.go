package department

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

type Repository interface {
	CreateDepartment(ctx context.Context, name, description string) (model.Department, error)
	GetDepartmentByID(ctx context.Context, id int) (model.Department, error)
	GetAllDepartments(ctx context.Context) ([]model.Department, error)
	UpdateDepartment(ctx context.Context, id int, name, description string) (model.Department, error)
	DeleteDepartment(ctx context.Context, id int) error
}

type impl struct {
	entClient *ent.Client
}

func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}