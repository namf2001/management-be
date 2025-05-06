package player

import (
	"management-be/internal/controller/player"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdatePlayerRequest represents the request body for updating a player
// @name UpdatePlayerRequest
type UpdatePlayerRequest struct {
	DepartmentID int    `json:"department_id" example:"1"`
	FullName     string `json:"full_name" example:"John Doe"`
	JerseyNumber int    `json:"jersey_number" example:"10"`
	Position     string `json:"position" example:"Forward"`
	DateOfBirth  string `json:"date_of_birth,omitempty" example:"1990-01-01"`
	HeightCm     int    `json:"height_cm,omitempty" example:"180"`
	WeightKg     int    `json:"weight_kg,omitempty" example:"75"`
	Phone        string `json:"phone,omitempty" example:"+1234567890"`
	Email        string `json:"email,omitempty" example:"john.doe@example.com"`
	IsActive     bool   `json:"is_active" example:"true"`
}

// UpdatePlayer
// @Summary      Update a player
// @Description  Update an existing player's information
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Player ID"
// @Param        player  body      UpdatePlayerRequest  true  "Player information"
// @Success      200  {object}  object{success=bool,data=PlayerResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      404  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/players/{id} [put]
func (h Handler) UpdatePlayer(ctx *gin.Context) {
	// Get player ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid player ID",
		})
		return
	}

	// Parse request body
	var req UpdatePlayerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	// Parse date of birth if provided
	var dob *time.Time
	if req.DateOfBirth != "" {
		parsedDOB, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid date format for date_of_birth (YYYY-MM-DD)",
			})
			return
		}
		dob = &parsedDOB
	}

	// Create input for controller
	input := player.InputPlayerController{
		DepartmentID: req.DepartmentID,
		FullName:     req.FullName,
		JerseyNumber: int32(req.JerseyNumber),
		Position:     req.Position,
		DateOfBirth:  dob,
		HeightCm:     int32(req.HeightCm),
		WeightKg:     int32(req.WeightKg),
		Phone:        req.Phone,
		Email:        req.Email,
		IsActive:     req.IsActive,
	}

	// Update player
	updatedPlayer, err := h.playerCtrl.UpdatePlayer(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Format date of birth if it exists
	var dateOfBirth string
	if updatedPlayer.DateOfBirth != nil && !updatedPlayer.DateOfBirth.IsZero() {
		dateOfBirth = updatedPlayer.DateOfBirth.Format("2006-01-02")
	}

	// Get jersey number, height and weight safely
	var jerseyNumber, heightCm, weightKg int
	if updatedPlayer.JerseyNumber != nil {
		jerseyNumber = int(*updatedPlayer.JerseyNumber)
	}
	if updatedPlayer.HeightCM != nil {
		heightCm = int(*updatedPlayer.HeightCM)
	}
	if updatedPlayer.WeightKG != nil {
		weightKg = int(*updatedPlayer.WeightKG)
	}

	// Return updated player
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": PlayerResponse{
			ID:           updatedPlayer.ID,
			DepartmentID: updatedPlayer.DepartmentID,
			FullName:     updatedPlayer.FullName,
			JerseyNumber: jerseyNumber,
			Position:     updatedPlayer.Position,
			DateOfBirth:  dateOfBirth,
			HeightCm:     heightCm,
			WeightKg:     weightKg,
			Phone:        updatedPlayer.Phone,
			Email:        updatedPlayer.Email,
			IsActive:     updatedPlayer.IsActive,
		},
	})
}
