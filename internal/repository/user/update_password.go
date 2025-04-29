package user

import (
	"context"

	"management-be/internal/repository/ent"

	pkgerrors "github.com/pkg/errors"
)

// UpdatePassword updates a user's password
func (i impl) UpdatePassword(ctx context.Context, userID int, hashedPassword string) error {
	_, err := i.entClient.User.UpdateOneID(userID).
		SetPassword(hashedPassword).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}
		return pkgerrors.WithStack(err)
	}
	return nil
}
