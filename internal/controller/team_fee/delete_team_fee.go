package team_fee

import (
	"context"
)

// DeleteTeamFee deletes a team fee by ID
func (i impl) DeleteTeamFee(ctx context.Context, id int) error {
	// First check if the fee exists
	_, err := i.repo.TeamFee().GetTeamFeeByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete the fee
	return i.repo.TeamFee().DeleteTeamFee(ctx, id)
}
