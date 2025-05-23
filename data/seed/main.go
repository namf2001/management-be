package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"management-be/internal/database"
	"management-be/internal/repository/ent"
)

func main() {
	log.Println("Starting database seeding...")

	// Initialize database connection
	db := database.New()
	client := db.Client()
	defer db.Close()

	ctx := context.Background()

	// Delete all existing data before seeding
	if err := deleteAllData(ctx, client); err != nil {
		log.Fatalf("Failed to delete existing data: %v", err)
	}

	// Seed data in the correct order to respect foreign key constraints
	if err := seedDepartments(ctx, client); err != nil {
		log.Fatalf("Failed to seed departments: %v", err)
	}

	if err := seedUsers(ctx, client); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	if err := seedTeams(ctx, client); err != nil {
		log.Fatalf("Failed to seed teams: %v", err)
	}

	if err := seedPlayers(ctx, client); err != nil {
		log.Fatalf("Failed to seed players: %v", err)
	}

	if err := seedMatches(ctx, client); err != nil {
		log.Fatalf("Failed to seed matches: %v", err)
	}

	if err := seedMatchPlayers(ctx, client); err != nil {
		log.Fatalf("Failed to seed match players: %v", err)
	}

	if err := seedTeamFees(ctx, client); err != nil {
		log.Fatalf("Failed to seed team fees: %v", err)
	}

	if err := seedPlayerStatistics(ctx, client); err != nil {
		log.Fatalf("Failed to seed player statistics: %v", err)
	}

	log.Println("Database seeding completed successfully!")
}

var tableSequences = []string{
	"departments",
	"users",
	"teams",
	"players",
	"matches",
	"match_players",
	"team_fees",
	"player_statistics",
}

func buildResetSequencesSQL() string {
	var statements []string
	for _, table := range tableSequences {
		statements = append(statements, fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", table))
	}
	return strings.Join(statements, ";\n") + ";"
}

func deleteAllData(ctx context.Context, client *ent.Client) error {
	log.Println("Deleting all existing data...")

	// Delete in reverse order of dependencies
	if _, err := client.PlayerStatistic.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.TeamFee.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.MatchPlayer.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.Match.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.Player.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.Team.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.User.Delete().Exec(ctx); err != nil {
		return err
	}

	if _, err := client.Department.Delete().Exec(ctx); err != nil {
		return err
	}

	// Reset all sequences to 1
	if _, err := client.ExecContext(ctx, buildResetSequencesSQL()); err != nil {
		return err
	}

	log.Println("Successfully deleted all existing data and reset ID sequences")
	return nil
}
