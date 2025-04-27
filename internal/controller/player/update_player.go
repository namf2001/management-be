package player

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/player"
	"time"
)

// UpdatePlayer updates an existing player in the database.
func (i impl) UpdatePlayer(ctx context.Context, id int, input InputPlayerController) (model.Player, error) {
	var dob time.Time
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

	player, err := i.repo.Player().UpdatePlayer(ctx, id, playerRepo)
	if err != nil {
		return model.Player{}, err
	}

	return player, nil
}
