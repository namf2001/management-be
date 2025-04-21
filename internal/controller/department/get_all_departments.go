package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetAllDepartments retrieves all departments
func (i impl) GetAllDepartments(ctx context.Context) ([]model.Department, error) {
	departmentRepo := i.repo.Department()
	departments, err := departmentRepo.GetAllDepartments(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(ErrDatabase)
	}

	return departments, nil
}
