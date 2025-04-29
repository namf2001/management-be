package player

import (
	"github.com/gin-gonic/gin"
	"management-be/internal/controller/player"
	"net/http"
	"time"
)

type CreatePlayerRequest struct {
	DepartmentID int    `json:"department_id" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	JerseyNumber int    `json:"jersey_number" binding:"required"`
	Position     string `json:"position" binding:"required"`
	DateOfBirth  string `json:"date_of_birth,omitempty"`
	HeightCm     int    `json:"height_cm,omitempty"`
	WeightKg     int    `json:"weight_kg,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
}

// CreatePlayer handles the request to create a new player
func (h Handler) CreatePlayer(ctx *gin.Context) {
	// Parse request body
	var req CreatePlayerRequest
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
		IsActive:     true, // New players are active by default
	}

	// Create player
	newPlayer, err := h.playerCtrl.CreatePlayer(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Format date of birth if it exists
	var dateOfBirth string
	if newPlayer.DateOfBirth != nil && !newPlayer.DateOfBirth.IsZero() {
		dateOfBirth = newPlayer.DateOfBirth.Format("2006-01-02")
	}

	// Get jersey number
	var jerseyNumber int
	if newPlayer.JerseyNumber != nil {
		jerseyNumber = int(*newPlayer.JerseyNumber)
	}

	// Get height and weight
	var heightCm, weightKg int
	if newPlayer.HeightCM != nil {
		heightCm = int(*newPlayer.HeightCM)
	}
	if newPlayer.WeightKG != nil {
		weightKg = int(*newPlayer.WeightKG)
	}

	// Get department info from controller if needed
	// In a real implementation, you would fetch the department name

	// Return created player
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": PlayerResponse{
			ID:           newPlayer.ID,
			DepartmentID: newPlayer.DepartmentID,
			// DepartmentName would be fetched from a department service
			DepartmentName: "", // This would be populated in a real implementation
			FullName:       newPlayer.FullName,
			JerseyNumber:   jerseyNumber,
			Position:       newPlayer.Position,
			DateOfBirth:    dateOfBirth,
			HeightCm:       heightCm,
			WeightKg:       weightKg,
			Phone:          newPlayer.Phone,
			Email:          newPlayer.Email,
			IsActive:       newPlayer.IsActive,
		},
	})
}
