package match

import (
	"context"
	"testing"
	"time"

	"management-be/internal/model"
	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"

	"github.com/stretchr/testify/require"
)

func TestCreateManyMatches(t *testing.T) {
	// Helper function to create test matches
	createTestMatches := func() []model.Match {
		matchDate1 := time.Date(2025, 6, 15, 15, 0, 0, 0, time.UTC)
		matchDate2 := time.Date(2025, 6, 22, 18, 0, 0, 0, time.UTC)
		
		return []model.Match{
			{
				OpponentTeamID: 1,
				MatchDate:      matchDate1,
				Venue:          "Batch Test Venue 1",
				IsHomeGame:     true,
				Status:         "scheduled",
				Notes:          "Batch Test Notes 1",
			},
			{
				OpponentTeamID: 2,
				MatchDate:      matchDate2,
				Venue:          "Batch Test Venue 2",
				IsHomeGame:     false,
				Status:         "scheduled",
				Notes:          "Batch Test Notes 2",
			},
		}
	}

	type args struct {
		matches []model.Match
		expErr  error
	}

	tcs := map[string]args{
		"success": {
			matches: createTestMatches(),
		},
		"success - empty list": {
			matches: []model.Match{},
		},
		"err - invalid opponent team ID": {
			matches: []model.Match{
				{
					OpponentTeamID: 999, // Non-existent team ID
					MatchDate:      time.Now().UTC(),
					Venue:          "Test Venue",
					IsHomeGame:     true,
					Status:         "scheduled",
					Notes:          "Test Notes",
				},
			},
			expErr: ErrDatabase,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Clean the tables first to avoid conflicts
				_, err := tx.ExecContext(context.Background(), "DELETE FROM matches")
				require.NoError(t, err)
				_, err = tx.ExecContext(context.Background(), "DELETE FROM teams")
				require.NoError(t, err)
				
				// Load team data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")

				repo := NewRepository(tx.Client())
				createdMatches, err := repo.CreateManyMatches(context.Background(), tc.matches)

				// then
				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Len(t, createdMatches, len(tc.matches))
					
					// If matches were created, verify they exist in the database
					if len(tc.matches) > 0 {
						for i, match := range createdMatches {
							require.NotZero(t, match.ID)
							require.Equal(t, tc.matches[i].OpponentTeamID, match.OpponentTeamID)
							
							// Compare time fields more carefully
							require.Equal(t, tc.matches[i].MatchDate.Year(), match.MatchDate.Year())
							require.Equal(t, tc.matches[i].MatchDate.Month(), match.MatchDate.Month())
							require.Equal(t, tc.matches[i].MatchDate.Day(), match.MatchDate.Day())
							require.Equal(t, tc.matches[i].MatchDate.Hour(), match.MatchDate.Hour())
							require.Equal(t, tc.matches[i].MatchDate.Minute(), match.MatchDate.Minute())
							
							require.Equal(t, tc.matches[i].Venue, match.Venue)
							require.Equal(t, tc.matches[i].IsHomeGame, match.IsHomeGame)
							require.Equal(t, tc.matches[i].Status, match.Status)
							require.Equal(t, tc.matches[i].Notes, match.Notes)

							// Verify the match was actually created in the database
							dbMatch, err := tx.Client().Match.Get(context.Background(), match.ID)
							require.NoError(t, err)
							require.NotNil(t, dbMatch)
						}
					}
				}
			})
		})
	}
}