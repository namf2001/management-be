package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/match"

	"time"
)

// UpdateMatchInput represents the input for updating a match
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

// UpdateMatch updates an existing match with transaction support
func (i impl) UpdateMatch(ctx context.Context, id int, input UpdateMatchInput) (model.Match, error) {
	var result model.Match

	// Check if match exists
	_, err := i.repo.Match().GetMatchByID(ctx, id)
	if err != nil {
		return model.Match{}, ErrMatchNotFound
	}

	// Execute the update operation within a transaction
	err = i.repo.WithTransaction(ctx, func(tx *ent.Tx) error {
		// Update match
		result, err = i.repo.Match().UpdateMatch(ctx, id, match.UpdateMatchInput{
			OpponentTeamID: input.OpponentTeamID,
			MatchDate:      input.MatchDate,
			Venue:          input.Venue,
			IsHomeGame:     input.IsHomeGame,
			OurScore:       input.OurScore,
			OpponentScore:  input.OpponentScore,
			Status:         input.Status,
			Notes:          input.Notes,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return model.Match{}, err
	}

	return result, nil
}
