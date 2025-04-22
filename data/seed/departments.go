package main

import (
	"context"
	"log"
	"time"

	"management-be/internal/repository/ent"
	"management-be/internal/repository/ent/department"

	"github.com/bxcodec/faker/v3"
)

// seedDepartments creates fake department data
func seedDepartments(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding departments...")

	// Define department names
	departmentNames := []string{
		"Engineering",
		"Marketing",
		"Sales",
		"Human Resources",
		"Finance",
		"Operations",
		"Research and Development",
		"Customer Support",
		"Legal",
		"Product Management",
	}

	// Create departments
	for _, name := range departmentNames {
		// Check if department already exists
		exists, err := client.Department.Query().Where(department.NameEQ(name)).Exist(ctx)
		if err != nil {
			return err
		}

		if exists {
			log.Printf("Department '%s' already exists, skipping", name)
			continue
		}

		description := faker.Sentence()
		now := time.Now()

		_, err = client.Department.Create().
			SetName(name).
			SetDescription(description).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d departments", len(departmentNames))
	return nil
}
