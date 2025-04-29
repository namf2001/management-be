package match

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent/match"
	"time"
)

// ListMatches retrieves matches based on the provided filters.
func (i impl) ListMatches(ctx context.Context, status string, startDate, endDate time.Time, opponentTeamID int) ([]model.Match, error) {
	query := i.entClient.Match.Query()

	if status != "" {
		query = query.Where(match.Status(status))
	}
	if !startDate.IsZero() {
		query = query.Where(match.MatchDateGTE(startDate))
	}
	if !endDate.IsZero() {
		query = query.Where(match.MatchDateLTE(endDate))
	}
	if opponentTeamID > 0 {
		query = query.Where(match.OpponentTeamID(opponentTeamID))
	}

	matches, err := query.All(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var result []model.Match
	for _, m := range matches {
		result = append(result, model.Match{
			ID:             m.ID,
			OpponentTeamID: m.OpponentTeamID,
			MatchDate:      m.MatchDate,
			Venue:          m.Venue,
			IsHomeGame:     m.IsHomeGame,
			OurScore:       m.OurScore,
			OpponentScore:  m.OpponentScore,
			Status:         m.Status,
			Notes:          m.Notes,
			CreatedAt:      m.CreatedAt,
			UpdatedAt:      m.UpdatedAt,
		})
	}

	return result, nil
}
