package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"management-be/internal/repository/ent"

	"github.com/bxcodec/faker/v3"
)

// seedTeamFees creates fake team fee data
func seedTeamFees(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding team fees...")

	// Number of fee records to create
	numFees := 20

	// Fee descriptions
	descriptions := []string{
		"Tournament registration fee",
		"Equipment purchase",
		"Uniform costs",
		"Training facility rental",
		"Coach payment",
		"Travel expenses",
		"Medical supplies",
		"Team building event",
		"League membership fee",
		"Insurance",
	}

	// Create team fees
	for i := 0; i < numFees; i++ {
		// Random amount between 100 and 5000
		amount := float64(rand.Intn(4901) + 100) // 100-5000
		
		// Random payment date (within the last 2 years)
		now := time.Now()
		daysAgo := rand.Intn(730) // 0-730 days (2 years)
		paymentDate := now.AddDate(0, 0, -daysAgo)
		
		// Random description
		var description string
		if rand.Intn(2) == 0 { // 50% chance of using predefined description
			description = descriptions[rand.Intn(len(descriptions))]
		} else {
			description = faker.Sentence() // 50% chance of using faker
		}
		
		_, err := client.TeamFee.Create().
			SetAmount(amount).
			SetPaymentDate(paymentDate).
			SetDescription(description).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d team fees", numFees)
	return nil
}