package department

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// GetDepartmentByID retrieves a department by its ID.
func (i impl) GetDepartmentByID(ctx context.Context, id int) (model.Department, error) {
	department, err := i.entClient.Department.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.Department{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.Department{}, pkgerrors.WithStack(err)
	}

	return model.Department{
		ID:          department.ID,
		Name:        department.Name,
		Description: department.Description,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
	}, nil
}
