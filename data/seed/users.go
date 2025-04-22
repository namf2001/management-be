package main

import (
	"context"
	"log"
	"time"

	"management-be/internal/repository/ent"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

// seedUsers creates fake user data
func seedUsers(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding users...")

	// Number of users to create
	numUsers := 10

	// Create users
	for i := 0; i < numUsers; i++ {
		username := faker.Username()
		email := faker.Email()
		fullName := faker.Name()
		
		// Generate a hashed password (using "password" as the default)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		
		now := time.Now()

		_, err = client.User.Create().
			SetUsername(username).
			SetEmail(email).
			SetFullName(fullName).
			SetPassword(string(hashedPassword)).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d users", numUsers)
	return nil
}