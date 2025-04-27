package model

import (
	"time"
)

// FootballMatch represents a football match from an external API
type FootballMatch struct {
	ID                int        `json:"id"`
	CompetitionName   string     `json:"competition_name"`
	SeasonStartDate   time.Time  `json:"season_start_date"`
	MatchDate         time.Time  `json:"match_date"`
	HomeTeamName      string     `json:"home_team_name"`
	HomeTeamShortName string     `json:"home_team_short_name"`
	HomeTeamLogo      string     `json:"home_team_logo"`
	AwayTeamName      string     `json:"away_team_name"`
	AwayTeamShortName string     `json:"away_team_short_name"`
	AwayTeamLogo      string     `json:"away_team_logo"`
	HomeScore         *int32     `json:"home_score,omitempty"`
	AwayScore         *int32     `json:"away_score,omitempty"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

// FootballMatchResponse represents the response from football-data.org API
type FootballMatchResponse struct {
	Competition struct {
		Name string `json:"name"`
	} `json:"competition"`
	Season struct {
		StartDate string `json:"startDate"`
	} `json:"season"`
	UtcDate  string       `json:"utcDate"`
	Status   string       `json:"status"`
	HomeTeam FootballTeam `json:"homeTeam"`
	AwayTeam FootballTeam `json:"awayTeam"`
	Score    Score        `json:"score"`
}

// FootballTeam represents a team in the football-data.org API response
type FootballTeam struct {
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Crest     string `json:"crest"`
}

// Score represents the score in the football-data.org API response
type Score struct {
	FullTime struct {
		Home *int `json:"home"`
		Away *int `json:"away"`
	} `json:"fullTime"`
}
