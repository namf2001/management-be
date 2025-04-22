package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"management-be/internal/repository/ent"

	"github.com/bxcodec/faker/v3"
)

// seedPlayers creates fake player data
func seedPlayers(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding players...")

	// Get all departments
	departments, err := client.Department.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(departments) == 0 {
		log.Println("No departments found. Please seed departments first.")
		return nil
	}

	// Player positions
	positions := []string{
		"Goalkeeper", "Defender", "Midfielder", "Forward", "Striker",
		"Center Back", "Left Back", "Right Back", "Defensive Midfielder",
		"Attacking Midfielder", "Winger", "Center Forward",
	}

	// Number of players to create
	numPlayers := 50

	// Create players
	for i := 0; i < numPlayers; i++ {
		fullName := faker.Name()
		position := positions[rand.Intn(len(positions))]
		jerseyNumber := int32(rand.Intn(99) + 1)

		// Random date of birth (18-40 years old)
		now := time.Now()
		yearsAgo := rand.Intn(22) + 18 // 18-40 years
		dob := now.AddDate(-yearsAgo, -rand.Intn(12), -rand.Intn(28))

		// Random height (165-195 cm)
		heightCm := int32(rand.Intn(30) + 165)

		// Random weight (60-95 kg)
		weightKg := int32(rand.Intn(35) + 60)

		// Random department
		departmentID := departments[rand.Intn(len(departments))].ID

		// Contact info
		phone := faker.Phonenumber()
		email := faker.Email()

		_, err := client.Player.Create().
			SetFullName(fullName).
			SetPosition(position).
			SetJerseyNumber(jerseyNumber).
			SetDateOfBirth(dob).
			SetHeightCm(heightCm).
			SetWeightKg(weightKg).
			SetDepartmentID(departmentID).
			SetPhone(phone).
			SetEmail(email).
			SetIsActive(true).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			// If there's a duplicate jersey number, try again with a different number
			if err.Error() == "ent: constraint failed: ERROR: duplicate key value violates unique constraint \"players_jersey_number_key\" (SQLSTATE 23505)" {
				log.Printf("Jersey number %d already taken, trying another number", jerseyNumber)
				i-- // retry this iteration
				continue
			}
			return err
		}
	}

	log.Printf("Successfully seeded %d players", numPlayers)
	return nil
}
