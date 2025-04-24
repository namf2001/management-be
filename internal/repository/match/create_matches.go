package match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// CreateManyMatches creates multiple matches at once.
// Note: The OnConflict().DoNothing() functionality is not directly supported by the ent framework.
// This implementation uses CreateBulk for efficient batch insertion.
func (i impl) CreateManyMatches(ctx context.Context, matches []model.Match) ([]model.Match, error) {
	now := time.Now()

	// Create a slice of match creators
	creators := make([]*ent.MatchCreate, len(matches))
	for idx, match := range matches {
		creators[idx] = i.entClient.Match.Create().
			SetOpponentTeamID(match.OpponentTeamID).
			SetMatchDate(match.MatchDate).
			SetVenue(match.Venue).
			SetIsHomeGame(match.IsHomeGame).
			SetStatus(match.Status).
			SetNotes(match.Notes).
			SetCreatedAt(now).
			SetUpdatedAt(now)

		// Set optional fields if they exist
		if match.OurScore != 0 {
			creators[idx].SetOurScore(match.OurScore)
		}
		if match.OpponentScore != 0 {
			creators[idx].SetOpponentScore(match.OpponentScore)
		}
	}

	// Use CreateBulk for efficient batch insertion
	entMatches, err := i.entClient.Match.CreateBulk(creators...).Save(ctx)
	if err != nil {
		return nil, err
	}

	// Convert ent.Match entities to model.Match
	result := make([]model.Match, len(entMatches))
	for idx, entMatch := range entMatches {
		result[idx] = model.Match{
			ID:             entMatch.ID,
			OpponentTeamID: entMatch.OpponentTeamID,
			MatchDate:      entMatch.MatchDate,
			Venue:          entMatch.Venue,
			IsHomeGame:     entMatch.IsHomeGame,
			OurScore:       entMatch.OurScore,
			OpponentScore:  entMatch.OpponentScore,
			Status:         entMatch.Status,
			Notes:          entMatch.Notes,
			CreatedAt:      entMatch.CreatedAt,
			UpdatedAt:      entMatch.UpdatedAt,
		}
	}

	return result, nil
}
