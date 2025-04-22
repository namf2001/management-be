package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"management-be/internal/repository/ent"
)

// seedMatchPlayers creates fake match player data
func seedMatchPlayers(ctx context.Context, client *ent.Client) error {
	log.Println("Seeding match players...")

	// Get all matches
	matches, err := client.Match.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		log.Println("No matches found. Please seed matches first.")
		return nil
	}

	// Get all players
	players, err := client.Player.Query().All(ctx)
	if err != nil {
		return err
	}

	if len(players) == 0 {
		log.Println("No players found. Please seed players first.")
		return nil
	}

	// Create match players
	for _, match := range matches {
		// Only create match players for completed or in_progress matches
		if match.Status != "completed" && match.Status != "in_progress" {
			continue
		}

		// Randomly select 11-18 players for this match
		numPlayers := rand.Intn(8) + 11 // 11-18 players
		
		// Shuffle players to get a random selection
		rand.Shuffle(len(players), func(i, j int) {
			players[i], players[j] = players[j], players[i]
		})
		
		// Use the first numPlayers from the shuffled list
		selectedPlayers := players[:numPlayers]
		
		now := time.Now()
		
		for _, player := range selectedPlayers {
			// Random minutes played (0-90)
			minutesPlayed := int32(rand.Intn(91))
			
			// Random goals (0-3)
			goalsScored := int32(rand.Intn(4))
			
			// Random assists (0-3)
			assists := int32(rand.Intn(4))
			
			// Random yellow cards (0-1)
			yellowCards := int32(rand.Intn(2))
			
			// Random red card (rare)
			redCard := rand.Intn(20) == 0 // 1/20 chance
			
			_, err := client.MatchPlayer.Create().
				SetMatchID(match.ID).
				SetPlayerID(player.ID).
				SetMinutesPlayed(minutesPlayed).
				SetGoalsScored(goalsScored).
				SetAssists(assists).
				SetYellowCards(yellowCards).
				SetRedCard(redCard).
				SetCreatedAt(now).
				SetUpdatedAt(now).
				Save(ctx)
				
			if err != nil {
				// If there's a unique constraint violation, skip this player for this match
				if err.Error() == "ent: constraint failed: UNIQUE constraint failed: match_players.match_id, match_players.player_id" {
					continue
				}
				return err
			}
		}
	}

	log.Println("Successfully seeded match players")
	return nil
}