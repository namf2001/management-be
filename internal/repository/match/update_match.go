package match

import (
	"context"
	"time"

	pkgerrors "github.com/pkg/errors"

	"management-be/internal/model"
	"management-be/internal/repository/ent"
)

type UpdateMatchInput struct {
	OpponentTeamID int       `json:"opponent_team_id"`
	MatchDate      time.Time `json:"match_date"`
	Venue          string    `json:"venue"`
	IsHomeGame     bool      `json:"is_home_game"`
	OurScore       int32     `json:"our_score"`
	OpponentScore  int32     `json:"opponent_score"`
	Status         string    `json:"status"`
	Notes          string    `json:"notes"`
}

// UpdateMatch updates an existing match with the provided details.
func (i impl) UpdateMatch(ctx context.Context, id int, input UpdateMatchInput) (model.Match, error) {
	now := time.Now()

	match, err := i.entClient.Match.UpdateOneID(id).
		SetOpponentTeamID(input.OpponentTeamID).
		SetMatchDate(input.MatchDate).
		SetVenue(input.Venue).
		SetIsHomeGame(input.IsHomeGame).
		SetOurScore(input.OurScore).
		SetOpponentScore(input.OpponentScore).
		SetStatus(input.Status).
		SetNotes(input.Notes).
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
