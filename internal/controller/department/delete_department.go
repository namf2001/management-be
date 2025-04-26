package department

import (
	"context"
	"management-be/internal/repository/ent"
)

// DeleteDepartment deletes a department by its ID
func (i impl) DeleteDepartment(ctx context.Context, id int) error {
	// Check if the department exists
	_, err := i.repo.Department().GetDepartmentByID(ctx, id)
	if err != nil {
		return err
	}

	// Execute the delete operation within a transaction
	err = i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		err = i.repo.Department().DeleteDepartment(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
