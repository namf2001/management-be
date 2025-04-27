package player

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/repository/ent"
)

// DeletePlayer deletes a player by ID
func (i impl) DeletePlayer(ctx context.Context, id int) error {
	err := i.entClient.Player.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}
	}

	return err
}
