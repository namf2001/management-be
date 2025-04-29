package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/match"
	"time"
)

// CreateMatchInput represents the input for creating a match
type CreateMatchInput struct {
	OpponentTeamID int       `json:"opponent_team_id"`
	MatchDate      time.Time `json:"match_date"`
	Venue          string    `json:"venue"`
	IsHomeGame     bool      `json:"is_home_game"`
	Notes          string    `json:"notes"`
}

// CreateMatch creates a new match
func (i impl) CreateMatch(ctx context.Context, input CreateMatchInput) (model.Match, error) {
	// Call the repository method
	match, err := i.repo.Match().CreateMatch(ctx, match.CreateMatchInput{
		OpponentTeamID: input.OpponentTeamID,
		MatchDate:      input.MatchDate,
		Venue:          input.Venue,
		IsHomeGame:     input.IsHomeGame,
		Notes:          input.Notes,
	})
	if err != nil {
		return model.Match{}, err
	}

	return match, nil
}
