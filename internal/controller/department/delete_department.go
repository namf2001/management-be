package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/repository/ent"
)

// DeleteDepartment deletes a department by its ID
func (i impl) DeleteDepartment(ctx context.Context, id int) error {
	// Execute the delete operation within a transaction
	err := i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		departmentRepo := i.repo.Department()

		// Check if department exists
		_, err := departmentRepo.GetDepartmentByID(ctx, id)
		if err != nil {
			return pkgerrors.WithStack(ErrDepartmentNotFound)
		}

		// Delete department
		err = departmentRepo.DeleteDepartment(ctx, id)
		if err != nil {
			return pkgerrors.WithStack(ErrDepartmentNotDeleted)
		}

		return nil
	})

	return err
}
