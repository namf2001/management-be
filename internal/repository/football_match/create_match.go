package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/matchesgateway"
	"time"
)

// CreateMatch saves a football match to the database
func (i impl) CreateMatch(ctx context.Context, match model.FootballMatch) (model.FootballMatch, error) {
	// Check if match already exists
	exists, err := i.entClient.MatchesGateway.
		Query().
		Where(
			matchesgateway.And(
				matchesgateway.CompetitionNameEQ(match.CompetitionName),
				matchesgateway.MatchDateEQ(match.MatchDate),
				matchesgateway.HomeTeamNameEQ(match.HomeTeamName),
				matchesgateway.AwayTeamNameEQ(match.AwayTeamName),
			),
		).
		Exist(ctx)

	if err != nil {
		return model.FootballMatch{}, err
	}

	if exists {
		// Find the match ID first
		existingMatch, err := i.entClient.MatchesGateway.
			Query().
			Where(
				matchesgateway.And(
					matchesgateway.CompetitionNameEQ(match.CompetitionName),
					matchesgateway.MatchDateEQ(match.MatchDate),
					matchesgateway.HomeTeamNameEQ(match.HomeTeamName),
					matchesgateway.AwayTeamNameEQ(match.AwayTeamName),
				),
			).
			Only(ctx)

		if err != nil {
			return model.FootballMatch{}, err
		}

		// Update the match
		updatedMatch, err := i.entClient.MatchesGateway.
			UpdateOneID(existingMatch.ID).
			SetNillableHomeScore(match.HomeScore).
			SetNillableAwayScore(match.AwayScore).
			SetStatus(match.Status).
			SetUpdatedAt(time.Now()).
			Save(ctx)

		if err != nil {
			return model.FootballMatch{}, err
		}

		return mapEntToModel(updatedMatch), nil
	}

	// Create new match
	now := time.Now()
	createdMatch, err := i.entClient.MatchesGateway.
		Create().
		SetCompetitionName(match.CompetitionName).
		SetSeasonStartDate(match.SeasonStartDate).
		SetMatchDate(match.MatchDate).
		SetHomeTeamName(match.HomeTeamName).
		SetHomeTeamShortName(match.HomeTeamShortName).
		SetHomeTeamLogo(match.HomeTeamLogo).
		SetAwayTeamName(match.AwayTeamName).
		SetAwayTeamShortName(match.AwayTeamShortName).
		SetAwayTeamLogo(match.AwayTeamLogo).
		SetNillableHomeScore(match.HomeScore).
		SetNillableAwayScore(match.AwayScore).
		SetStatus(match.Status).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return model.FootballMatch{}, err
	}

	return mapEntToModel(createdMatch), nil
}
