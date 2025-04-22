package player

import (
	"context"
	"management-be/internal/model"
	"time"
)

// CreatePlayer creates a new player in the database.
func (i impl) CreatePlayer(ctx context.Context, departmentID int, fullName, position string, jerseyNumber int32, dateOfBirth *time.Time, heightCm, weightKg int32, phone, email string, isActive bool) (model.Player, error) {
	//TODO implement me
	panic("implement me")
}
