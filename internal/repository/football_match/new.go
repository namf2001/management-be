package football_match

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

// Repository defines the interface for football match repository operations
type Repository interface {
	// CreateMatch saves a football match to the database
	CreateMatch(ctx context.Context, match model.FootballMatch) (model.FootballMatch, error)
	// CreateMatches saves multiple football matches to the database
	CreateMatches(ctx context.Context, matches []model.FootballMatch) ([]model.FootballMatch, error)
	// GetMatchesByCompetition gets football matches by competition name
	GetMatchesByCompetition(ctx context.Context, competitionName string) ([]model.FootballMatch, error)
	// GetMatchesByDateRange gets football matches within a date range
	GetMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatch, error)
	// GetMatchesByStatus gets football matches by status
	GetMatchesByStatus(ctx context.Context, status string) ([]model.FootballMatch, error)
	// GetPreviousDayMatches gets football matches from the previous day
	GetPreviousDayMatches(ctx context.Context) ([]model.FootballMatch, error)
}

type impl struct {
	entClient *ent.Client
}

// NewRepository creates a new football match repository
func NewRepository(entClient *ent.Client) Repository {
	return &impl{
		entClient: entClient,
	}
}

// Helper function to map ent.MatchesGateway to model.FootballMatch
func mapEntToModel(match *ent.MatchesGateway) model.FootballMatch {
	var homeScore, awayScore *int32
	if match.HomeScore != 0 {
		homeScore = &match.HomeScore
	}
	if match.AwayScore != 0 {
		awayScore = &match.AwayScore
	}

	var deletedAt *time.Time
	if !match.DeletedAt.IsZero() {
		deletedAt = &match.DeletedAt
	}

	return model.FootballMatch{
		ID:                match.ID,
		CompetitionName:   match.CompetitionName,
		SeasonStartDate:   match.SeasonStartDate,
		MatchDate:         match.MatchDate,
		HomeTeamName:      match.HomeTeamName,
		HomeTeamShortName: match.HomeTeamShortName,
		HomeTeamLogo:      match.HomeTeamLogo,
		AwayTeamName:      match.AwayTeamName,
		AwayTeamShortName: match.AwayTeamShortName,
		AwayTeamLogo:      match.AwayTeamLogo,
		HomeScore:         homeScore,
		AwayScore:         awayScore,
		Status:            match.Status,
		CreatedAt:         match.CreatedAt,
		UpdatedAt:         match.UpdatedAt,
		DeletedAt:         deletedAt,
	}
}

// Helper function to map []*ent.MatchesGateway to []model.FootballMatch
func mapEntsToModels(matches []*ent.MatchesGateway) []model.FootballMatch {
	result := make([]model.FootballMatch, len(matches))
	for i, match := range matches {
		result[i] = mapEntToModel(match)
	}
	return result
}
