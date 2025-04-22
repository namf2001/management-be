package main

import (
	"context"
	"log"
	"management-be/internal/database"
)

func main() {
	log.Println("Starting database seeding...")

	// Initialize database connection
	db := database.New()
	client := db.Client()
	defer db.Close()

	ctx := context.Background()

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
