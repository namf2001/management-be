package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetTeamStatistics handles the request to get statistics for a team
func (h Handler) GetTeamStatistics(ctx *gin.Context) {
	// Get team ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team ID",
		})
		return
	}

	// Get team statistics from controller
	stats, err := h.teamCtrl.GetTeamStatistics(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Team not found or statistics unavailable",
		})
		return
	}

	// Prepare recent form data (last 5 matches)
	recentForm := make([]gin.H, 0)
	for i, match := range stats.Matches {
		if i >= 5 {
			break
		}

		var result string
		if match.OurScore > match.OpponentScore {
			result = "W"
		} else if match.OurScore < match.OpponentScore {
			result = "L"
		} else {
			result = "D"
		}

		recentForm = append(recentForm, gin.H{
			"match_date": match.MatchDate.Format(time.RFC3339),
			"result":     result,
			"score":      strconv.Itoa(match.OurScore) + "-" + strconv.Itoa(match.OpponentScore),
		})
	}

	// Calculate goals scored and conceded
	goalsScored := 0
	goalsConceded := 0
	homeMatches := 0
	homeWins := 0
	homeLosses := 0
	homeDraws := 0
	awayMatches := 0
	awayWins := 0
	awayLosses := 0
	awayDraws := 0

	for _, match := range stats.Matches {
		goalsScored += match.OurScore
		goalsConceded += match.OpponentScore

		// Determine home/away performance
		if match.Venue == "home" {
			homeMatches++
			if match.OurScore > match.OpponentScore {
				homeWins++
			} else if match.OurScore < match.OpponentScore {
				homeLosses++
			} else {
				homeDraws++
			}
		} else {
			awayMatches++
			if match.OurScore > match.OpponentScore {
				awayWins++
			} else if match.OurScore < match.OpponentScore {
				awayLosses++
			} else {
				awayDraws++
			}
		}
	}

	// Return team statistics
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"match_summary": gin.H{
				"total_matches":  stats.TotalMatches,
				"wins":           stats.Wins,
				"losses":         stats.Losses,
				"draws":          stats.Draws,
				"goals_scored":   goalsScored,
				"goals_conceded": goalsConceded,
			},
			"recent_form": recentForm,
			"performance_by_venue": gin.H{
				"home": gin.H{
					"matches": homeMatches,
					"wins":    homeWins,
					"losses":  homeLosses,
					"draws":   homeDraws,
				},
				"away": gin.H{
					"matches": awayMatches,
					"wins":    awayWins,
					"losses":  awayLosses,
					"draws":   awayDraws,
				},
			},
		},
	})
}