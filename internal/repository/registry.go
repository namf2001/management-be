package repository

import (
	"context"
	"management-be/internal/repository/department"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/player"
	"management-be/internal/repository/team"
	"management-be/internal/repository/user"
)

type Registry interface {
	User() user.Repository
	Department() department.Repository
	Team() team.Repository
	WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error
}

type impl struct {
	entConn    *ent.Client
	user       user.Repository
	department department.Repository
	team       team.Repository
	player     player.Repository
}

func NewRegistry(entConn *ent.Client) Registry {
	return &impl{
		entConn:    entConn,
		user:       user.NewRepository(entConn),
		department: department.NewRepository(entConn),
		player:     player.NewRepository(entConn),
		team:       team.NewRepository(entConn),
	}
}

func (i *impl) User() user.Repository {
	return i.user
}

func (i *impl) Department() department.Repository {
	return i.department
}

func (i *impl) Team() team.Repository {
	return i.team
}

func (i *impl) Player() player.Repository {
	return i.player
}

// WithTransaction executes the given function within a transaction
func (i *impl) WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := i.entConn.Tx(ctx)
	if err != nil {
		return err
	}

	// Execute the function
	if err := fn(tx); err != nil {
		// Roll back the transaction in case of error
		if rerr := tx.Rollback(); rerr != nil {
			// Return both errors if rollback fails
			return rerr
		}
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
