package player

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/matchplayer"
	"management-be/internal/repository/ent/playerstatistic"
)

// DeletePlayer deletes a player by ID
func (i impl) DeletePlayer(ctx context.Context, id int) error {
	// First delete related match_players records
	_, err := i.entClient.MatchPlayer.Delete().
		Where(matchplayer.PlayerID(id)).
		Exec(ctx)
	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	// Delete related player_statistics records
	_, err = i.entClient.PlayerStatistic.Delete().
		Where(playerstatistic.PlayerID(id)).
		Exec(ctx)
	if err != nil {
		return pkgerrors.WithStack(ErrDatabase)
	}

	err = i.entClient.Player.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return pkgerrors.WithStack(ErrNotFound)
		}
		return pkgerrors.WithStack(ErrDatabase)
	}

	return nil
}
