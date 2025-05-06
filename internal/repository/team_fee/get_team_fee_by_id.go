package team_fee

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"management-be/internal/repository/ent/teamfee"
)

// GetTeamFeeByID retrieves a team fee by its ID
func (i impl) GetTeamFeeByID(ctx context.Context, id int) (model.TeamFee, error) {
	teamFee, err := i.entClient.TeamFee.Query().
		Where(teamfee.ID(id)).
		Where(teamfee.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return model.TeamFee{}, pkgerrors.WithStack(err)
	}

	return model.TeamFee{
		ID:          teamFee.ID,
		Amount:      teamFee.Amount,
		PaymentDate: teamFee.PaymentDate,
		Description: teamFee.Description,
		CreatedAt:   teamFee.CreatedAt,
		UpdatedAt:   teamFee.UpdatedAt,
	}, nil
}
