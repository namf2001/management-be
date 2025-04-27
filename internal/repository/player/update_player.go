package player

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent"
	"time"
)

func (i impl) UpdatePlayer(ctx context.Context, id int, input InputPlayer) (model.Player, error) {
	now := time.Now()

	// Update player using ent client
	player, err := i.entClient.Player.UpdateOneID(id).
		SetFullName(input.FullName).
		SetPosition(input.Position).
		SetJerseyNumber(input.JerseyNumber).
		SetDateOfBirth(input.DateOfBirth).
		SetHeightCm(input.HeightCm).
		SetWeightKg(input.WeightKg).
		SetPhone(input.Phone).
		SetEmail(input.Email).
		SetIsActive(input.IsActive).
		SetUpdatedAt(now).
		SetDepartmentID(input.DepartmentID).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return model.Player{}, pkgerrors.WithStack(ErrNotFound)
		}
		return model.Player{}, err
	}

	// Convert ent.Player to model.Player
	return model.Player{
		ID:           player.ID,
		DepartmentID: player.DepartmentID,
		FullName:     player.FullName,
		Position:     player.Position,
		JerseyNumber: &player.JerseyNumber,
		DateOfBirth:  &player.DateOfBirth,
		HeightCM:     &player.HeightCm,
		WeightKG:     &player.WeightKg,
		Phone:        player.Phone,
		Email:        player.Email,
		IsActive:     player.IsActive,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}, nil
}
