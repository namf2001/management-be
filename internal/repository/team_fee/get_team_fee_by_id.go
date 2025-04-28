package team_fee

import (
	"context"
	"management-be/internal/model"
	"management-be/internal/repository/ent/teamfee"
	"time"
)

// GetTeamFee retrieves a team fee by its ID
func (i impl) GetTeamFee(ctx context.Context, id int) (model.TeamFee, error) {
	teamFee, err := i.entClient.TeamFee.Query().
		Where(teamfee.ID(id)).
		Only(ctx)
	if err != nil {
		return model.TeamFee{}, err
	}

	var deletedAt *time.Time
	if !teamFee.DeletedAt.IsZero() {
		deletedAt = &teamFee.DeletedAt
	}

	return model.TeamFee{
		ID:          teamFee.ID,
		Amount:      teamFee.Amount,
		PaymentDate: teamFee.PaymentDate,
		Description: teamFee.Description,
		CreatedAt:   teamFee.CreatedAt,
		UpdatedAt:   teamFee.UpdatedAt,
		DeletedAt:   deletedAt,
	}, nil
}
