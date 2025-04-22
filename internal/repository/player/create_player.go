package player

import (
	"context"
	"management-be/internal/model"
	"time"
)

func (i impl) CreatePlayer(ctx context.Context, departmentID int, fullName, position string, jerseyNumber int32, dateOfBirth time.Time, heightCm, weightKg int32, phone, email string, isActive bool) (model.Player, error) {
	// Create player using ent client
	player, err := i.entClient.Player.Create().
		SetFullName(fullName).
		SetPosition(position).
		SetJerseyNumber(jerseyNumber).
		SetDateOfBirth(dateOfBirth).
		SetHeightCm(heightCm).
		SetWeightKg(weightKg).
		SetPhone(phone).
		SetEmail(email).
		SetIsActive(isActive).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetDepartmentID(departmentID).
		Save(ctx)

	if err != nil {
		return model.Player{}, err
	}

	// Convert ent.Player to model.Player
	return model.Player{
		ID:           player.ID,
		DepartmentID: player.DepartmentID,
		FullName:     player.FullName,
		Position:     player.Position,
		JerseyNumber: player.JerseyNumber,
		DateOfBirth:  &player.DateOfBirth,
		HeightCm:     player.HeightCm,
		WeightKg:     player.WeightKg,
		Phone:        player.Phone,
		Email:        player.Email,
		IsActive:     player.IsActive,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}, nil
}
