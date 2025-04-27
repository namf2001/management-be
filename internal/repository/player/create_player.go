package player

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

type InputPlayer struct {
	DepartmentID int
	FullName     string
	Position     string
	JerseyNumber int32
	DateOfBirth  time.Time
	HeightCm     int32
	WeightKg     int32
	Phone        string
	Email        string
	IsActive     bool
}

// CreatePlayer creates a new player with the provided details.
func (i impl) CreatePlayer(ctx context.Context, input InputPlayer) (model.Player, error) {
	// Create player using ent client
	player, err := i.entClient.Player.Create().
		SetFullName(input.FullName).
		SetPosition(input.Position).
		SetJerseyNumber(input.JerseyNumber).
		SetDateOfBirth(input.DateOfBirth).
		SetHeightCm(input.HeightCm).
		SetWeightKg(input.WeightKg).
		SetPhone(input.Phone).
		SetEmail(input.Email).
		SetIsActive(input.IsActive).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetDepartmentID(input.DepartmentID).
		Save(ctx)

	if err != nil {
		return model.Player{}, pkgerrors.WithStack(ErrDatabase)
	}

	var heightCM, weightKG *int32
	if player.HeightCm != 0 {
		heightCM = &player.HeightCm
	}
	if player.WeightKg != 0 {
		weightKG = &player.WeightKg
	}

	// Convert ent.Player to model.Player
	return model.Player{
		ID:           player.ID,
		DepartmentID: player.DepartmentID,
		FullName:     player.FullName,
		Position:     player.Position,
		JerseyNumber: &player.JerseyNumber,
		DateOfBirth:  &player.DateOfBirth,
		HeightCM:     heightCM,
		WeightKG:     weightKG,
		Phone:        player.Phone,
		Email:        player.Email,
		IsActive:     player.IsActive,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}, nil
}
