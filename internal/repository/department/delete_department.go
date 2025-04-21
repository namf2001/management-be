package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/repository/ent"
)

// DeleteDepartment deletes a department by its ID.
func (i impl) DeleteDepartment(ctx context.Context, id int) error {
	// Check if department exists
	_, err := i.entClient.Department.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}
		return pkgerrors.WithStack(ErrDatabase)
	}

	// Delete department
	err = i.entClient.Department.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	return nil
}