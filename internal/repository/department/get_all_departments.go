package department

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetAllDepartments retrieves all departments from the database.
func (i impl) GetAllDepartments(ctx context.Context) ([]model.Department, error) {
	departments, err := i.entClient.Department.Query().All(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(ErrDatabase)
	}

	result := make([]model.Department, len(departments))
	for idx, department := range departments {
		result[idx] = model.Department{
			ID:          department.ID,
			Name:        department.Name,
			Description: department.Description,
			CreatedAt:   department.CreatedAt,
			UpdatedAt:   department.UpdatedAt,
		}
	}

	return result, nil
}