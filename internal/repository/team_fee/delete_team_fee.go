package team_fee

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/repository/ent"
)

// DeleteTeamFee deletes a team fee by ID
func (i impl) DeleteTeamFee(ctx context.Context, id int) error {
	err := i.entClient.TeamFee.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}
		return pkgerrors.WithStack(err)
	}

	return nil
}
