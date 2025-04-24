package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchplayer"
	"time"
)

// UpdateMatchPlayers updates the players for a given match.
// This implementation first deletes all existing match players for the match
// and then creates new ones based on the provided players slice.
func (i impl) UpdateMatchPlayers(ctx context.Context, matchID int, players []model.MatchPlayer) error {
	// First, delete all existing match players for this match
	_, err := i.entClient.MatchPlayer.Delete().
		Where(matchplayer.MatchID(matchID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Then create new match players
	now := time.Now()
	for _, player := range players {
		_, err := i.entClient.MatchPlayer.Create().
			SetMatchID(matchID).
			SetPlayerID(player.PlayerID).
			SetMinutesPlayed(int32(player.MinutesPlayed)).
			SetGoalsScored(int32(player.GoalsScored)).
			SetAssists(int32(player.Assists)).
			SetYellowCards(int32(player.YellowCards)).
			SetRedCard(player.RedCard).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
