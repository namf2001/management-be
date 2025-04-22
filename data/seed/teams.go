package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"management-be/internal/repository/ent"

	"github.com/bxcodec/faker/v3"
)

// seedTeams creates fake team data
func seedTeams(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding teams...")

	// Number of teams to create
	numTeams := 15

	// Create teams
	for i := 0; i < numTeams; i++ {
		name := faker.Word() + " " + faker.Word() + " FC " + fmt.Sprintf("%d", i+1)
		companyName := faker.Word() + " " + faker.Word() + " Inc"
		contactPerson := faker.Name()
		contactPhone := faker.Phonenumber()
		contactEmail := faker.Email()
		now := time.Now()

		_, err := client.Team.Create().
			SetName(name).
			SetCompanyName(companyName).
			SetContactPerson(contactPerson).
			SetContactPhone(contactPhone).
			SetContactEmail(contactEmail).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d teams", numTeams)
	return nil
}
