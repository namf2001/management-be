package scheduler

import (
	"context"
	"log"
	"management-be/internal/controller/football_match"
	"time"
)

// ScheduleFootballMatchFetching schedules the football match data fetching to run at the beginning of each day
func ScheduleFootballMatchFetching(controller football_match.Controller) {
	// Run immediately on startup
	go fetchFootballMatches(controller)

	// Schedule to run at the beginning of each day (00:05 AM)
	go func() {
		for {
			now := time.Now()
			// Calculate time until next run (00:05 AM tomorrow)
			nextRun := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 5, 0, 0, now.Location())
			duration := nextRun.Sub(now)

			log.Printf("Football match data fetching scheduled to run in %v", duration)

			// Sleep until next run time
			time.Sleep(duration)

			// Fetch football matches
			fetchFootballMatches(controller)
		}
	}()
}

// fetchFootballMatches fetches football matches from the previous day
func fetchFootballMatches(controller football_match.Controller) {
	log.Println("Fetching football matches from the previous day...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	err := controller.FetchAndSavePreviousDayMatches(ctx)
	if err != nil {
		log.Printf("Error fetching football matches: %v", err)
		return
	}

	log.Println("Football matches fetched and saved successfully")
}
