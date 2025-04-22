package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// CreateDepartment creates a new department in the database.
func (i impl) CreateDepartment(ctx context.Context, name, description string) (model.Department, error) {
	newDepartment := i.entClient.Department.Create().
		SetName(name).
		SetDescription(description).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now())

	createdDepartment, err := newDepartment.Save(ctx)
	if err != nil {
		return model.Department{}, pkgerrors.WithStack(ErrDatabase)
	}

	return model.Department{
		ID:          createdDepartment.ID,
		Name:        createdDepartment.Name,
		Description: createdDepartment.Description,
		CreatedAt:   createdDepartment.CreatedAt,
		UpdatedAt:   createdDepartment.UpdatedAt,
	}, nil
}
