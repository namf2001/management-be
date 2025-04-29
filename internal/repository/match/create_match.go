package match

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// CreateMatch is creating a match with the provided details.
func (i impl) CreateMatch(ctx context.Context, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, notes string) (model.Match, error) {
	now := time.Now()

	match, err := i.entClient.Match.Create().
		SetOpponentTeamID(opponentTeamID).
		SetMatchDate(matchDate).
		SetVenue(venue).
		SetIsHomeGame(isHomeGame).
		SetStatus("scheduled").
		SetNotes(notes).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return model.Match{}, pkgerrors.WithStack(ErrDatabase)
	}

	return model.Match{
		ID:             match.ID,
		OpponentTeamID: match.OpponentTeamID,
		MatchDate:      match.MatchDate,
		Venue:          match.Venue,
		IsHomeGame:     match.IsHomeGame,
		Status:         match.Status,
		Notes:          match.Notes,
		CreatedAt:      match.CreatedAt,
		UpdatedAt:      match.UpdatedAt,
	}, nil
}
