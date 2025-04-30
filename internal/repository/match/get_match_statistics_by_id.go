package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/match"
	"management-be/internal/repository/ent/matchplayer"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error) {
	// First check if the match exists
	exists, err := i.entClient.Match.Query().
		Where(match.ID(matchID)).
		Exist(ctx)
	if err != nil {
		return model.MatchStatistics{}, pkgerrors.WithStack(err)
	}
	if !exists {
		return model.MatchStatistics{}, pkgerrors.WithStack(ErrNotFound)
	}

	// Get match players
	players, err := i.entClient.MatchPlayer.Query().
		Where(matchplayer.MatchID(matchID)).
		WithPlayer().
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.MatchStatistics{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.MatchStatistics{}, pkgerrors.WithStack(err)
	}

	var stats model.MatchStatistics
	stats.MatchSummary.TotalPlayers = int32(len(players))

	for _, p := range players {
		stats.MatchSummary.TotalMinutesPlayed += p.MinutesPlayed
		stats.MatchSummary.TotalGoals += p.GoalsScored
		stats.MatchSummary.TotalAssists += p.Assists
		stats.MatchSummary.TotalYellowCards += p.YellowCards
		if p.RedCard {
			stats.MatchSummary.TotalRedCards++
		}

		stats.PlayerPerformance = append(stats.PlayerPerformance, struct {
			PlayerID      int    `json:"player_id"`
			PlayerName    string `json:"player_name"`
			Position      string `json:"position"`
			MinutesPlayed int32  `json:"minutes_played"`
			GoalsScored   int32  `json:"goals_scored"`
			Assists       int32  `json:"assists"`
			YellowCards   int32  `json:"yellow_cards"`
			RedCard       bool   `json:"red_card"`
		}{
			PlayerID:      p.PlayerID,
			PlayerName:    p.Edges.Player.FullName,
			Position:      p.Edges.Player.Position,
			MinutesPlayed: p.MinutesPlayed,
			GoalsScored:   p.GoalsScored,
			Assists:       p.Assists,
			YellowCards:   p.YellowCards,
			RedCard:       p.RedCard,
		})
	}

	return stats, nil
}
