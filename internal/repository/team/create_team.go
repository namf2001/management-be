package team

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"management-be/internal/model"
	"time"
)

// CreateTeam creates a new team in the database
func (i impl) CreateTeam(ctx context.Context, name, companyName, contactPerson, contactPhone, contactEmail string) (model.Team, error) {
	now := time.Now()
	team, err := i.entClient.Team.
		Create().
		SetName(name).
		SetCompanyName(companyName).
		SetContactPerson(contactPerson).
		SetContactPhone(contactPhone).
		SetContactEmail(contactEmail).
		SetCreatedAt(now).
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