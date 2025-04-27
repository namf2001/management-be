package football_match

import (
	"context"
	"fmt"
	"management-be/internal/model"
	"time"
)

// FetchAndSavePreviousDayMatches fetches matches from the previous day and saves them to the database
func (i *impl) FetchAndSavePreviousDayMatches(ctx context.Context) error {
	// Fetch matches from the previous day
	matches, err := i.gateway.FetchPreviousDayMatches(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch previous day matches: %w", err)
	}

	// Convert API response to model
	footballMatches := make([]model.FootballMatch, 0, len(matches))
	for _, match := range matches {
		// Parse season start date
		seasonStartDate, err := time.Parse("2006-01-02", match.Season.StartDate)
		if err != nil {
			return fmt.Errorf("failed to parse season start date: %w", err)
		}

		// Parse match date
		matchDate, err := time.Parse(time.RFC3339, match.UtcDate)
		if err != nil {
			return fmt.Errorf("failed to parse match date: %w", err)
		}

		// Create football match model
		footballMatch := model.FootballMatch{
			CompetitionName:   match.Competition.Name,
			SeasonStartDate:   seasonStartDate,
			MatchDate:         matchDate,
			HomeTeamName:      match.HomeTeam.Name,
			HomeTeamShortName: match.HomeTeam.ShortName,
			HomeTeamLogo:      match.HomeTeam.Crest,
			AwayTeamName:      match.AwayTeam.Name,
			AwayTeamShortName: match.AwayTeam.ShortName,
			AwayTeamLogo:      match.AwayTeam.Crest,
			Status:            match.Status,
		}

		// Add scores if available
		if match.Score.FullTime.Home != nil && match.Score.FullTime.Away != nil {
			homeScore := int32(*match.Score.FullTime.Home)
			awayScore := int32(*match.Score.FullTime.Away)
			footballMatch.HomeScore = &homeScore
			footballMatch.AwayScore = &awayScore
		}

		footballMatches = append(footballMatches, footballMatch)
	}

	// Save matches to the database
	_, err = i.repo.FootballMatch().CreateMatches(ctx, footballMatches)
	if err != nil {
		return fmt.Errorf("failed to save matches to database: %w", err)
	}

	return nil
}
