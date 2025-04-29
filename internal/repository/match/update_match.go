package match

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// UpdateMatch updates an existing match with the provided details.
func (i impl) UpdateMatch(ctx context.Context, id int, opponentTeamID int, matchDate time.Time, venue string, isHomeGame bool, ourScore, opponentScore int32, status, notes string) (model.Match, error) {
	now := time.Now()

	match, err := i.entClient.Match.UpdateOneID(id).
		SetOpponentTeamID(opponentTeamID).
		SetMatchDate(matchDate).
		SetVenue(venue).
		SetIsHomeGame(isHomeGame).
		SetOurScore(ourScore).
		SetOpponentScore(opponentScore).
		SetStatus(status).
		SetNotes(notes).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return model.Match{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.Match{}, pkgerrors.WithStack(err)
	}

	return model.Match{
		ID:             match.ID,
		OpponentTeamID: match.OpponentTeamID,
		MatchDate:      match.MatchDate,
		Venue:          match.Venue,
		IsHomeGame:     match.IsHomeGame,
		OurScore:       match.OurScore,
		OpponentScore:  match.OpponentScore,
		Status:         match.Status,
		Notes:          match.Notes,
		CreatedAt:      match.CreatedAt,
		UpdatedAt:      match.UpdatedAt,
	}, nil
}
