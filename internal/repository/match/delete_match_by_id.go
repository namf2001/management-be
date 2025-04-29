package match

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/repository/ent"
)

// DeleteMatchByID deletes a match by its ID.
func (i impl) DeleteMatchByID(ctx context.Context, id int) error {
	err := i.entClient.Match.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}

		return pkgerrors.WithStack(err)
	}

	return nil
}
