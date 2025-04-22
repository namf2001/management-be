package team

import (
	"context"
	"entgo.io/ent/dialect/sql"
	pkgerrors "github.com/pkg/errors"
)

// DeleteTeam deletes a team from the database
func (i impl) DeleteTeam(ctx context.Context, id int) error {
	// Check if team exists
	exists, err := i.entClient.Team.Query().
		Where(sql.FieldEQ("id", id)).
		Exist(ctx)

	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	if !exists {
		return pkgerrors.WithStack(ErrNotFound)
	}

	// Delete team
	err = i.entClient.Team.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	return nil
}