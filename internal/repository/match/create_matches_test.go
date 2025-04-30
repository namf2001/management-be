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

// TestCreateManyMatches tests the CreateManyMatches function
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
		matches   []model.Match
		expResult []model.Match
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			matches:   createTestMatches(),
			expResult: createTestMatches(),
		},
		"success - empty list": {
			matches:   []model.Match{},
			expResult: []model.Match{},
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
				// Load team data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_team.sql")

				repo := NewRepository(tx.Client())
				createdMatches, err := repo.CreateManyMatches(context.Background(), tc.matches)

				// then
				if tc.expErr != nil {
					require.Error(t, err)
					if s == "err - invalid opponent team ID" {
						// For this specific test case, we expect a database foreign key error
						require.Contains(t, err.Error(), "foreign key constraint")
					} else {
						require.ErrorIs(t, err, tc.expErr)
					}
				} else {
					require.NoError(t, err)
					require.Len(t, createdMatches, len(tc.matches))

					// If matches were created, verify they exist in the database
					if len(tc.matches) > 0 {
						for i, match := range createdMatches {
							require.NotZero(t, match.ID)
							require.Equal(t, tc.matches[i].OpponentTeamID, match.OpponentTeamID)
							require.Equal(t, tc.matches[i].MatchDate, match.MatchDate)
							require.Equal(t, tc.matches[i].Venue, match.Venue)
							require.Equal(t, tc.matches[i].IsHomeGame, match.IsHomeGame)
							require.Equal(t, tc.matches[i].Status, match.Status)
							require.Equal(t, tc.matches[i].Notes, match.Notes)
							require.NotZero(t, match.CreatedAt)
							require.NotZero(t, match.UpdatedAt)

							// For model.Match, DeletedAt is a pointer that should be nil or point to a zero time
							if match.DeletedAt != nil {
								require.True(t, match.DeletedAt.IsZero())
							}
						}
					}
				}
			})
		})
	}
}
