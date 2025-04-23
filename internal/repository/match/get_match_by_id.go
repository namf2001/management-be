package match

import (
	"context"
	"management-be/internal/model"
)

// GetMatch retrieves a match by its ID.
func (i impl) GetMatch(ctx context.Context, id int) (model.Match, error) {
	match, err := i.entClient.Match.Get(ctx, id)
	if err != nil {
		return model.Match{}, err
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
