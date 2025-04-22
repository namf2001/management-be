package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"management-be/internal/repository/ent"

	"github.com/bxcodec/faker/v3"
)

// seedMatches creates fake match data
func seedMatches(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding matches...")

	// Get all teams
	teams, err := client.Team.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(teams) < 2 {
		log.Println("Not enough teams found. Please seed at least 2 teams first.")
		return nil
	}

	// Match statuses
	statuses := []string{"scheduled", "in_progress", "completed", "cancelled", "postponed"}

	// Venues
	venues := []string{
		"City Stadium", "Central Park", "Sports Complex", "University Field",
		"Community Center", "Olympic Stadium", "Memorial Field", "Downtown Arena",
	}

	// Number of matches to create
	numMatches := 30

	// Create matches
	for i := 0; i < numMatches; i++ {
		// Random opponent team
		opponentTeamIndex := rand.Intn(len(teams))
		opponentTeam := teams[opponentTeamIndex]

		// Random match date (between 6 months ago and 6 months in the future)
		now := time.Now()
		daysOffset := rand.Intn(365) - 182 // -182 to +182 days
		matchDate := now.AddDate(0, 0, daysOffset)

		// Random venue
		venue := venues[rand.Intn(len(venues))]

		// Random home/away status
		isHomeGame := rand.Intn(2) == 1

		// Random status
		status := statuses[rand.Intn(len(statuses))]

		// Scores (only if match is completed)
		var ourScore, opponentScore *int32
		if status == "completed" {
			ourScoreVal := int32(rand.Intn(5))
			opponentScoreVal := int32(rand.Intn(5))
			ourScore = &ourScoreVal
			opponentScore = &opponentScoreVal
		}

		// Random notes
		notes := ""
		if rand.Intn(3) == 0 { // 1/3 chance of having notes
			notes = faker.Sentence()
		}

		matchCreate := client.Match.Create().
			SetOpponentTeamID(opponentTeam.ID).
			SetMatchDate(matchDate).
			SetVenue(venue).
			SetIsHomeGame(isHomeGame).
			SetStatus(status).
			SetNotes(notes).
			SetCreatedAt(now).
			SetUpdatedAt(now)

		// Set scores if available
		if ourScore != nil {
			matchCreate.SetOurScore(*ourScore)
		}
		if opponentScore != nil {
			matchCreate.SetOpponentScore(*opponentScore)
		}

		_, err := matchCreate.Save(ctx)
		if err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d matches", numMatches)
	return nil
}
