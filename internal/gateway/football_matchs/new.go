package football_matchs

import (
	"context"
	"encoding/json"
	"fmt"
	"management-be/internal/model"
	"net/http"
	"time"
)

// Gateway defines the interface for football matches gateway operations
type Gateway interface {
	// FetchMatchesByCompetition fetches matches for a specific competition
	FetchMatchesByCompetition(ctx context.Context, competitionCode string) ([]model.FootballMatchResponse, error)

	// FetchMatchesByDateRange fetches matches within a date range
	FetchMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatchResponse, error)

	// FetchTopLeaguesMatches fetches matches from the top 5 leagues
	FetchTopLeaguesMatches(ctx context.Context) ([]model.FootballMatchResponse, error)

	// FetchPreviousDayMatches fetches matches from the previous day for the top 5 leagues
	FetchPreviousDayMatches(ctx context.Context) ([]model.FootballMatchResponse, error)
}

type impl struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewGateway creates a new football matches gateway
func NewGateway(apiKey, baseURL string) Gateway {
	return &impl{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Top 5 European football leagues competition codes
const (
	PremierLeague = "PL"  // English Premier League
	LaLiga        = "PD"  // Spanish La Liga
	Bundesliga    = "BL1" // German Bundesliga
	SerieA        = "SA"  // Italian Serie A
	Ligue1        = "FL1" // French Ligue 1
)

// FetchMatchesByCompetition fetches matches for a specific competition
func (g *impl) FetchMatchesByCompetition(ctx context.Context, competitionCode string) ([]model.FootballMatchResponse, error) {
	url := fmt.Sprintf("%s/competitions/%s/matches", g.baseURL, competitionCode)
	return g.fetchMatches(ctx, url)
}

// FetchMatchesByDateRange fetches matches within a date range
func (g *impl) FetchMatchesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.FootballMatchResponse, error) {
	dateFrom := startDate.Format("2006-01-02")
	dateTo := endDate.Format("2006-01-02")
	url := fmt.Sprintf("%s/matches?dateFrom=%s&dateTo=%s", g.baseURL, dateFrom, dateTo)
	return g.fetchMatches(ctx, url)
}

// FetchTopLeaguesMatches fetches matches from the top 5 leagues
func (g *impl) FetchTopLeaguesMatches(ctx context.Context) ([]model.FootballMatchResponse, error) {
	var allMatches []model.FootballMatchResponse

	leagues := []string{PremierLeague, LaLiga, Bundesliga, SerieA, Ligue1}

	for _, league := range leagues {
		matches, err := g.FetchMatchesByCompetition(ctx, league)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch matches for league %s: %w", league, err)
		}
		allMatches = append(allMatches, matches...)
	}

	return allMatches, nil
}

// FetchPreviousDayMatches fetches matches from the previous day for the top 5 leagues
func (g *impl) FetchPreviousDayMatches(ctx context.Context) ([]model.FootballMatchResponse, error) {
	// Calculate yesterday's date
	yesterday := time.Now().AddDate(0, 0, -1)
	// Format as YYYY-MM-DD
	yesterdayStr := yesterday.Format("2006-01-02")

	// Create URL to fetch matches for yesterday
	url := fmt.Sprintf("%s/matches?dateFrom=%s&dateTo=%s", g.baseURL, yesterdayStr, yesterdayStr)

	// Fetch all matches from yesterday
	allMatches, err := g.fetchMatches(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch previous day matches: %w", err)
	}

	// Filter matches for top 5 leagues
	topLeagueMatches := make([]model.FootballMatchResponse, 0)
	topLeagues := map[string]bool{
		"Premier League":   true,
		"Primera Division": true,
		"Bundesliga":       true,
		"Serie A":          true,
		"Ligue 1":          true,
	}

	for _, match := range allMatches {
		if topLeagues[match.Competition.Name] {
			topLeagueMatches = append(topLeagueMatches, match)
		}
	}

	return topLeagueMatches, nil
}

// fetchMatches is a helper function to fetch matches from the API
func (g *impl) fetchMatches(ctx context.Context, url string) ([]model.FootballMatchResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Auth-Token", g.apiKey)

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var response struct {
		Matches []model.FootballMatchResponse `json:"matches"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Matches, nil
}
