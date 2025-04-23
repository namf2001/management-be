package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchplayer"
)

func (i impl) GetMatchStatistics(ctx context.Context, matchID int) (model.MatchStatistics, error) {
	// Get match players
	players, err := i.entClient.MatchPlayer.Query().
		Where(matchplayer.MatchID(matchID)).
		WithPlayer().
		All(ctx)
	if err != nil {
		return model.MatchStatistics{}, err
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
