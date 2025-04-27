package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/player"
	"time"
)

// InputPlayerController is the input for the CreatePlayer function.
type InputPlayerController struct {
	DepartmentID int
	FullName     string
	Position     string
	JerseyNumber int32
	DateOfBirth  *time.Time
	HeightCm     int32
	WeightKg     int32
	Phone        string
	Email        string
	IsActive     bool
}

// CreatePlayer creates a new player in the database.
func (i impl) CreatePlayer(ctx context.Context, input InputPlayerController) (model.Player, error) {
	var dob time.Time

	_, err := i.repo.Department().GetDepartmentByID(ctx, input.DepartmentID)
	if err != nil {
		return model.Player{}, err
	}

	if input.DateOfBirth != nil {
		dob = *input.DateOfBirth
	}

	playerRepo := player.InputPlayer{
		DepartmentID: input.DepartmentID,
		FullName:     input.FullName,
		Position:     input.Position,
		JerseyNumber: input.JerseyNumber,
		DateOfBirth:  dob,
		HeightCm:     input.HeightCm,
		WeightKg:     input.WeightKg,
		Phone:        input.Phone,
		Email:        input.Email,
		IsActive:     input.IsActive,
	}

	player, err := i.repo.Player().CreatePlayer(ctx, playerRepo)
	if err != nil {
		return model.Player{}, err
	}

	return player, nil
}
