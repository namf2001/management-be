package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"management-be/internal/repository/ent"
)

// seedPlayerStatistics creates player statistics with random data
func seedPlayerStatistics(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding player statistics...")

	// Get all players
	players, err := client.Player.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(players) == 0 {
		log.Println("No players found. Please seed players first.")
		return nil
	}

	now := time.Now()

	// For each player, generate random statistics
	for _, player := range players {
		// Generate random statistics
		totalMatches := int32(rand.Intn(30) + 1)  // 1-30 matches
		totalMinutesPlayed := totalMatches * int32(rand.Intn(60) + 30)  // 30-90 minutes per match
		totalGoals := int32(rand.Intn(20))  // 0-19 goals
		totalAssists := int32(rand.Intn(15))  // 0-14 assists
		totalYellowCards := int32(rand.Intn(10))  // 0-9 yellow cards
		totalRedCards := int32(rand.Intn(3))  // 0-2 red cards

		// Create player statistics
		_, err = client.PlayerStatistic.Create().
			SetPlayerID(player.ID).
			SetTotalMatches(totalMatches).
			SetTotalMinutesPlayed(totalMinutesPlayed).
			SetTotalGoals(totalGoals).
			SetTotalAssists(totalAssists).
			SetTotalYellowCards(totalYellowCards).
			SetTotalRedCards(totalRedCards).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(ctx)

		if err != nil {
			// If there's a unique constraint violation, skip this player
			if err.Error() == "ent: constraint failed: UNIQUE constraint failed: player_statistics.player_id" {
				continue
			}
			return err
		}
	}

	log.Printf("Successfully seeded statistics for %d players", len(players))
	return nil
}
