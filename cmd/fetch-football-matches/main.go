package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"management-be/internal/controller/football_match"
	"management-be/internal/gateway/football_matchs"
	"management-be/internal/model"
	"management-be/internal/repository"
	"management-be/internal/repository/ent"
	"os"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Define command-line flags
	dateStr := flag.String("date", "", "Date to fetch matches for (YYYY-MM-DD format)")
	usePrevDay := flag.Bool("prev-day", true, "Fetch matches from the previous day")
	flag.Parse()

	// Set up context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Get API key from environment variable or use the provided one
	apiKey := os.Getenv("FOOTBALL_DATA_KEY")
	if apiKey == "" {
		apiKey = "3e51e2d6acab4d3a89ea86bc3f112fdd" // Use the provided key if not set in environment
	}

	// Initialize database connection
	fmt.Println("Connecting to database...")
	dbConn, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbConn.Close()

	// Create ent client
	drv := entsql.OpenDB(dialect.Postgres, dbConn)
	entClient := ent.NewClient(ent.Driver(drv))
	defer entClient.Close()

	// Initialize repository registry
	repo := repository.NewRegistry(entClient)

	// Initialize football matches gateway
	gateway := football_matchs.NewGateway(apiKey, "https://api.football-data.org/v4")

	// Initialize football match controller
	controller := football_match.NewController(repo, gateway)

	var matches []model.FootballMatch

	if *dateStr != "" {
		// Parse the provided date
		date, err := time.Parse("2006-01-02", *dateStr)
		if err != nil {
			log.Fatalf("Invalid date format. Please use YYYY-MM-DD: %v", err)
		}

		// Fetch matches for the specified date
		fmt.Printf("Fetching matches for date: %s\n", date.Format("2006-01-02"))

		// Fetch matches from API
		apiMatches, err := gateway.FetchMatchesByDateRange(ctx, date, date)
		if err != nil {
			log.Fatalf("Failed to fetch matches: %v", err)
		}

		// Convert API response to model
		footballMatches := convertAPIResponseToModel(apiMatches)

		// Save matches to database
		savedMatches, err := repo.FootballMatch().CreateMatches(ctx, footballMatches)
		if err != nil {
			log.Fatalf("Failed to save matches: %v", err)
		}

		fmt.Printf("Successfully saved %d matches for date %s\n", len(savedMatches), date.Format("2006-01-02"))

		// Get matches from database for the specified date
		startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
		matches, err = controller.GetMatchesByDateRange(ctx, startOfDay, endOfDay)
		if err != nil {
			log.Fatalf("Failed to retrieve matches: %v", err)
		}
	} else if *usePrevDay {
		// Fetch and save previous day matches
		fmt.Println("Fetching and saving previous day matches...")
		err = controller.FetchAndSavePreviousDayMatches(ctx)
		if err != nil {
			log.Fatalf("Failed to fetch and save previous day matches: %v", err)
		}

		fmt.Println("Successfully fetched and saved previous day matches!")

		// Retrieve the saved matches
		fmt.Println("\nRetrieving saved matches...")
		matches, err = controller.GetPreviousDayMatches(ctx)
		if err != nil {
			log.Fatalf("Failed to retrieve previous day matches: %v", err)
		}
	}

	// Display the matches
	fmt.Printf("Found %d matches:\n", len(matches))
	for i, match := range matches {
		homeScore := "-"
		awayScore := "-"
		if match.HomeScore != nil {
			homeScore = fmt.Sprintf("%d", *match.HomeScore)
		}
		if match.AwayScore != nil {
			awayScore = fmt.Sprintf("%d", *match.AwayScore)
		}

		fmt.Printf("%d. %s: %s %s-%s %s (%s)\n",
			i+1,
			match.CompetitionName,
			match.HomeTeamName,
			homeScore,
			awayScore,
			match.AwayTeamName,
			match.Status,
		)

		// Print more details for debugging
		fmt.Printf("   Match date: %s\n", match.MatchDate.Format("2006-01-02 15:04:05"))
		fmt.Printf("   Home team: %s (%s)\n", match.HomeTeamName, match.HomeTeamShortName)
		fmt.Printf("   Away team: %s (%s)\n", match.AwayTeamName, match.AwayTeamShortName)
		if match.HomeScore != nil {
			fmt.Printf("   Home score: %d\n", *match.HomeScore)
		} else {
			fmt.Printf("   Home score: nil\n")
		}
		if match.AwayScore != nil {
			fmt.Printf("   Away score: %d\n", *match.AwayScore)
		} else {
			fmt.Printf("   Away score: nil\n")
		}
		fmt.Printf("   Status: %s\n", match.Status)
		fmt.Println()
	}
}

func initDB() (*sql.DB, error) {
	// Get database connection parameters from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "management-football")
	password := getEnv("DB_PASSWORD", "management-football")
	dbname := getEnv("DB_NAME", "management-football")
	schema := getEnv("DB_SCHEMA", "public")

	// Create database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		user, password, host, port, dbname, schema)

	// Connect to database
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// convertAPIResponseToModel converts the API response to the model format
func convertAPIResponseToModel(matches []model.FootballMatchResponse) []model.FootballMatch {
	footballMatches := make([]model.FootballMatch, 0, len(matches))

	for _, match := range matches {
		// Parse season start date
		seasonStartDate, err := time.Parse("2006-01-02", match.Season.StartDate)
		if err != nil {
			log.Printf("Warning: Failed to parse season start date for match: %v", err)
			// Use current date as fallback
			seasonStartDate = time.Now()
		}

		// Parse match date
		matchDate, err := time.Parse(time.RFC3339, match.UtcDate)
		if err != nil {
			log.Printf("Warning: Failed to parse match date for match: %v", err)
			// Skip this match if we can't parse the date
			continue
		}

		// Create football match model
		now := time.Now()
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
			CreatedAt:         now,
			UpdatedAt:         now,
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

	return footballMatches
}
