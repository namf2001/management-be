package team

import (
	"context"
	"entgo.io/ent/dialect/sql"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
)

// GetTeamByID retrieves a team by its ID
func (i impl) GetTeamByID(ctx context.Context, id int) (model.Team, error) {
	team, err := i.entClient.Team.Query().
		Where(sql.FieldEQ("id", id)).
		Only(ctx)

	if err != nil {
		return model.Team{}, pkgerrors.WithStack(ErrDatabase)
	}

	return model.Team{
		ID:            team.ID,
		Name:          team.Name,
		CompanyName:   team.CompanyName,
		ContactPerson: team.ContactPerson,
		ContactPhone:  team.ContactPhone,
		ContactEmail:  team.ContactEmail,
		CreatedAt:     team.CreatedAt,
		UpdatedAt:     team.UpdatedAt,
	}, nil
}