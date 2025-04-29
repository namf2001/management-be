package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/pkg/testent"
	"management-be/internal/repository/ent"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetAllPlayers(t *testing.T) {
	type args struct {
		filter      FilterGetAllPlayers
		expResults  []model.Player
		expTotal    int
		expectEmpty bool
		expErr      error
	}

	// Create test player data
	jerseyNum1 := int32(10)
	jerseyNum2 := int32(7)
	dob1 := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	dob2 := time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
	height1 := int32(180)
	height2 := int32(175)
	weight1 := int32(75)
	weight2 := int32(70)

	player1 := model.Player{
		ID:           1,
		DepartmentID: 1,
		FullName:     "Test Player",
		Position:     "Forward",
		JerseyNumber: &jerseyNum1,
		DateOfBirth:  &dob1,
		HeightCM:     &height1,
		WeightKG:     &weight1,
		Phone:        "1234567890",
		Email:        "test@example.com",
		IsActive:     true,
	}

	player2 := model.Player{
		ID:           2,
		DepartmentID: 2,
		FullName:     "Test Defender",
		Position:     "Defender",
		JerseyNumber: &jerseyNum2,
		DateOfBirth:  &dob2,
		HeightCM:     &height2,
		WeightKG:     &weight2,
		Phone:        "0987654321",
		Email:        "defender@example.com",
		IsActive:     true,
	}

	tcs := map[string]args{
		"success - get all players": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     1,
			},
			expResults: []model.Player{player1, player2},
			expTotal:   2,
		},
		"success - empty page": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     2,
			},
			expResults:  []model.Player{},
			expTotal:    2,
			expectEmpty: true,
		},
		"success - filter by department": {
			filter: FilterGetAllPlayers{
				PageSize:     10,
				Page:         1,
				DepartmentID: ptrInt(1),
			},
			expResults: []model.Player{player1},
			expTotal:   1,
		},
		"success - filter by active status": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     1,
				IsActive: ptrBool(true),
			},
			expResults: []model.Player{player1, player2},
			expTotal:   2,
		},
		"success - filter by position": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     1,
				Position: "Defender",
			},
			expResults: []model.Player{player2},
			expTotal:   1,
		},
		"success - invalid page number (defaults to page 1)": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     0,
			},
			expResults: []model.Player{player1, player2},
			expTotal:   2,
		},
		"success - negative page number (defaults to page 1)": {
			filter: FilterGetAllPlayers{
				PageSize: 10,
				Page:     -1,
			},
			expResults: []model.Player{player1, player2},
			expTotal:   2,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testent.WithEntTx(t, func(tx *ent.Tx) {
				// Load test data
				testent.LoadTestSQLFile(t, tx, "testdata/insert_department.sql")
				testent.LoadTestSQLFile(t, tx, "testdata/insert_player.sql")

				repo := NewRepository(tx.Client())
				players, total, err := repo.GetAllPlayers(
					context.Background(),
					tc.filter,
				)

				if tc.expErr != nil {
					require.Error(t, err)
					require.EqualError(t, err, tc.expErr.Error())
					return
				}

				require.NoError(t, err)
				require.Equal(t, tc.expTotal, total)

				if tc.expectEmpty {
					require.Empty(t, players)
				} else {
					require.Len(t, players, len(tc.expResults))

					for i, expected := range tc.expResults {
						actual := players[i]
						require.Equal(t, expected.ID, actual.ID)
						require.Equal(t, expected.DepartmentID, actual.DepartmentID)
						require.Equal(t, expected.FullName, actual.FullName)
						require.Equal(t, expected.Position, actual.Position)

						if expected.JerseyNumber != nil {
							require.NotNil(t, actual.JerseyNumber)
							require.Equal(t, *expected.JerseyNumber, *actual.JerseyNumber)
						}

						if expected.DateOfBirth != nil {
							require.NotNil(t, actual.DateOfBirth)
							require.Equal(t, expected.DateOfBirth.UTC(), actual.DateOfBirth.UTC())
						}

						if expected.HeightCM != nil {
							require.NotNil(t, actual.HeightCM)
							require.Equal(t, *expected.HeightCM, *actual.HeightCM)
						}

						if expected.WeightKG != nil {
							require.NotNil(t, actual.WeightKG)
							require.Equal(t, *expected.WeightKG, *actual.WeightKG)
						}

						require.Equal(t, expected.Phone, actual.Phone)
						require.Equal(t, expected.Email, actual.Email)
						require.Equal(t, expected.IsActive, actual.IsActive)
					}
				}
			})
		})
	}
}

// Helper functions to create pointers
// Helper functions to create pointers
func ptrInt(i int) *int {
	return &i
}

func ptrBool(b bool) *bool {
	return &b
}
