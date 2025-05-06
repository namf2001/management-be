package match_statistics

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/matchplayer"
)

func (i impl) GetMatchPlayers(ctx context.Context, matchID int) ([]model.MatchPlayer, error) {
	players, err := i.entClient.MatchPlayer.Query().
		Where(matchplayer.MatchID(matchID)).
		WithPlayer().
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, pkgerrors.WithStack(ErrNotFound)
		}
		return nil, pkgerrors.WithStack(err)
	}

	var matchPlayers []model.MatchPlayer
	for _, player := range players {
		matchPlayers = append(matchPlayers, model.MatchPlayer{
			ID:            player.ID,
			PlayerID:      player.PlayerID,
			MatchID:       player.MatchID,
			MinutesPlayed: player.MinutesPlayed,
			GoalsScored:   player.GoalsScored,
			Assists:       player.Assists,
			YellowCards:   player.YellowCards,
			RedCard:       player.RedCard,
			CreatedAt:     player.CreatedAt,
			UpdatedAt:     player.UpdatedAt,
		})
	}

	return matchPlayers, nil
}
