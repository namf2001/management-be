package player

import (
	"management-be/internal/controller/player"
	"management-be/internal/pkg/unit"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePlayerRequest represents the request body for creating a player
// @name CreatePlayerRequest
type CreatePlayerRequest struct {
	DepartmentID int    `json:"department_id" validate:"required,min=1" example:"1"`
	FullName     string `json:"full_name" validate:"required,min=2,max=100" example:"John Doe"`
	JerseyNumber int    `json:"jersey_number" validate:"required,min=1,max=999" example:"10"`
	Position     string `json:"position" validate:"required" example:"Forward"`
	DateOfBirth  string `json:"date_of_birth" validate:"omitempty,datetime=2006-01-02" example:"1990-01-01"`
	HeightCm     int    `json:"height_cm" validate:"omitempty,min=50,max=300" example:"180"`
	WeightKg     int    `json:"weight_kg" validate:"omitempty,min=30,max=200" example:"75"`
	Phone        string `json:"phone" validate:"omitempty,min=5,max=20" example:"+1234567890"`
	Email        string `json:"email" validate:"omitempty,email" example:"john.doe@example.com"`
}

// CreatePlayer
// @Summary      Create a new player
// @Description  Create a new player with their details
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        player  body      CreatePlayerRequest  true  "Player information"
// @Success      201  {object}  object{success=bool,data=PlayerResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/players [post]
func (h Handler) CreatePlayer(ctx *gin.Context) {
	// Parse request body
	var req CreatePlayerRequest
	if !unit.ValidateJSON(ctx, &req) {
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
