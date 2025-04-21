package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetDepartmentByID retrieves a department by its ID
func (i impl) GetDepartmentByID(ctx context.Context, id int) (model.Department, error) {
	departmentRepo := i.repo.Department()
	department, err := departmentRepo.GetDepartmentByID(ctx, id)
	if err != nil {
		return model.Department{}, pkgerrors.WithStack(ErrDepartmentNotFoundByID)
	}

	return department, nil
}
