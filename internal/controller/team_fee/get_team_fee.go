package team_fee

import (
	"context"
	"management-be/internal/model"
)

// GetTeamFee retrieves a team fee by its ID
func (i impl) GetTeamFee(ctx context.Context, id int) (model.TeamFee, error) {
	return i.repo.TeamFee().GetTeamFeeByID(ctx, id)
}
