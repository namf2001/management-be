package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// UpdateDepartment updates an existing department
func (i impl) UpdateDepartment(ctx context.Context, id int, name, description string) (model.Department, error) {
	var department model.Department

	// Execute the update operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		departmentRepo := i.repo.Department()

		// Check if department exists
		_, err := departmentRepo.GetDepartmentByID(ctx, id)
		if err != nil {
			return pkgerrors.WithStack(ErrDepartmentNotFound)
		}

		// Update department
		updatedDepartment, err := departmentRepo.UpdateDepartment(ctx, id, name, description)
		if err != nil {
			return pkgerrors.WithStack(ErrDepartmentNotUpdated)
		}

		department = updatedDepartment
		return nil
	})

	if err != nil {
		return model.Department{}, err
	}

	return department, nil
}
