package player

import (
	"context"
	"database/sql"
	"management-be/internal/model"
	"management-be/internal/repository/player"
)

// UpdatePlayer updates an existing player in the database.
func (i impl) UpdatePlayer(ctx context.Context, id int, input InputPlayerController) (model.Player, error) {
	var dbDOB sql.NullTime
	if input.DateOfBirth != nil && !input.DateOfBirth.IsZero() {
		dbDOB = sql.NullTime{Time: *input.DateOfBirth, Valid: true}
	} else {
		dbDOB = sql.NullTime{Valid: false}
	}

	playerRepo := player.InputPlayer{
		DepartmentID: input.DepartmentID,
		FullName:     input.FullName,
		Position:     input.Position,
		JerseyNumber: input.JerseyNumber,
		DateOfBirth:  dbDOB.Time,
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
