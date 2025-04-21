package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// CreateDepartment creates a new department
func (i impl) CreateDepartment(ctx context.Context, name, description string) (model.Department, error) {
	var department model.Department

	// Execute the create operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		// Create the department
		departmentRepo := i.repo.Department()
		createdDepartment, err := departmentRepo.CreateDepartment(ctx, name, description)
		if err != nil {
			return pkgerrors.WithStack(ErrDepartmentNotCreated)
		}

		department = createdDepartment
		return nil
	})

	if err != nil {
		return model.Department{}, err
	}

	return department, nil
}
