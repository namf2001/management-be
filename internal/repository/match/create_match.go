package match

import (
	"context"
	"management-be/internal/model"
	"time"

	pkgerrors "github.com/pkg/errors"
)

type CreateMatchInput struct {
	OpponentTeamID int       `json:"opponent_team_id"`
	MatchDate      time.Time `json:"match_date"`
	Venue          string    `json:"venue"`
	IsHomeGame     bool      `json:"is_home_game"`
	Notes          string    `json:"notes"`
}

// CreateMatch is creating a match with the provided details.
func (i impl) CreateMatch(ctx context.Context, input CreateMatchInput) (model.Match, error) {
	// Validate match date is not empty
	if input.MatchDate.IsZero() {
		return model.Match{}, pkgerrors.WithStack(ErrDatabase)
	}

	now := time.Now()
	match, err := i.entClient.Match.Create().
		SetOpponentTeamID(input.OpponentTeamID).
		SetMatchDate(input.MatchDate).
		SetVenue(input.Venue).
		SetIsHomeGame(input.IsHomeGame).
		SetStatus("scheduled").
		SetNotes(input.Notes).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return model.Match{}, pkgerrors.WithStack(err)
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
