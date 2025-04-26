package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
)

// DeleteDepartment deletes a department by its ID.
func (i impl) DeleteDepartment(ctx context.Context, id int) error {
	// Delete department
	err := i.entClient.Department.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	return nil
}
