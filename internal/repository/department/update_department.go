package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// UpdateDepartment updates an existing department in the database.
func (i impl) UpdateDepartment(ctx context.Context, id int, name, description string) (model.Department, error) {
	// Check if department exists
	_, err := i.entClient.Department.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.Department{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.Department{}, pkgerrors.WithStack(err)
	}

	// Update department
	updatedDepartment, err := i.entClient.Department.UpdateOneID(id).
		SetName(name).
		SetDescription(description).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return model.Department{}, pkgerrors.WithStack(err)
	}

	return model.Department{
		ID:          updatedDepartment.ID,
		Name:        updatedDepartment.Name,
		Description: updatedDepartment.Description,
		CreatedAt:   updatedDepartment.CreatedAt,
		UpdatedAt:   updatedDepartment.UpdatedAt,
	}, nil
}
