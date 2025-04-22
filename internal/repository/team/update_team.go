package team

import (
	"context"
	"entgo.io/ent/dialect/sql"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// UpdateTeam updates an existing team in the database
func (i impl) UpdateTeam(ctx context.Context, id int, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error) {
	// Check if team exists
	exists, err := i.entClient.Team.Query().
		Where(sql.FieldEQ("id", id)).
		Exist(ctx)

	if err != nil {
		return model.Team{}, pkgerrors.WithStack(ErrDatabase)
	}

	if !exists {
		return model.Team{}, pkgerrors.WithStack(ErrNotFound)
	}

	// Update team
	now := time.Now()
	team, err := i.entClient.Team.
		UpdateOneID(id).
		SetName(name).
		SetCompanyName(companyName).
		SetContactPerson(contactPerson).
		SetContactPhone(contactPhone).
		SetContactEmail(contactEmail).
		SetUpdatedAt(now).
		Save(ctx)

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