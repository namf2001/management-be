package repository

import (
	"context"
	"management-be/internal/repository/department"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/user"
)

type Registry interface {
	User() user.Repository
	Department() department.Repository
	WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error
}

type impl struct {
	entConn    *ent.Client
	user       user.Repository
	department department.Repository
}

func NewRegistry(entConn *ent.Client) Registry {
	return &impl{
		entConn:    entConn,
		user:       user.NewRepository(entConn),
		department: department.NewRepository(entConn),
	}
}

func (i *impl) User() user.Repository {
	return i.user
}

func (i *impl) Department() department.Repository {
	return i.department
}

// WithTransaction executes the given function within a transaction
func (i *impl) WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := i.entConn.Tx(ctx)
	if err != nil {
		return err
	}

	// Execute the function
	if err := fn(tx); err != nil {
		// Rollback the transaction in case of error
		if rerr := tx.Rollback(); rerr != nil {
			// Return both errors if rollback fails
			return rerr
		}
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
