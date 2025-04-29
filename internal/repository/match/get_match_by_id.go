package match

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

// GetMatchByID retrieves a match by its ID.
func (i impl) GetMatchByID(ctx context.Context, id int) (model.Match, error) {
	match, err := i.entClient.Match.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.Match{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.Match{}, pkgerrors.WithStack(ErrDatabase)
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
