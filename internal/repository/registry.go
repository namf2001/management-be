package repository

import (
	"context"
	"management-be/internal/repository/department"
	"management-be/internal/repository/ent"
	"management-be/internal/repository/match"
	"management-be/internal/repository/match_statistics"
	"management-be/internal/repository/player"
	"management-be/internal/repository/player_statistics"
	"management-be/internal/repository/team"
	"management-be/internal/repository/team_fee"
	"management-be/internal/repository/user"
)

type Registry interface {
	User() user.Repository
	Department() department.Repository
	Team() team.Repository
	Player() player.Repository
	Match() match.Repository
	TeamFee() team_fee.Repository
	PlayerStatistics() player_statistics.Repository
	MatchStatistics() match_statistics.Repository
	WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error
}

type impl struct {
	entConn          *ent.Client
	user             user.Repository
	department       department.Repository
	team             team.Repository
	player           player.Repository
	match            match.Repository
	teamFee          team_fee.Repository
	playerStatistics player_statistics.Repository
	matchStatistics  match_statistics.Repository
}

func NewRegistry(entConn *ent.Client) Registry {
	return &impl{
		entConn:          entConn,
		user:             user.NewRepository(entConn),
		department:       department.NewRepository(entConn),
		player:           player.NewRepository(entConn),
		team:             team.NewRepository(entConn),
		match:            match.NewRepository(entConn),
		teamFee:          team_fee.NewRepository(entConn),
		playerStatistics: player_statistics.NewRepository(entConn),
		matchStatistics:  match_statistics.NewRepository(entConn),
	}
}

func (i *impl) User() user.Repository {
	return i.user
}

func (i *impl) Department() department.Repository {
	return i.department
}

func (i *impl) Team() team.Repository {
	return i.team
}

func (i *impl) Player() player.Repository {
	return i.player
}

func (i *impl) PlayerStatistics() player_statistics.Repository {
	return i.playerStatistics
}

func (i *impl) Match() match.Repository {
	return i.match
}

func (i *impl) TeamFee() team_fee.Repository {
	return i.teamFee
}

func (i *impl) MatchStatistics() match_statistics.Repository {
	return i.matchStatistics
}

// WithTransaction executes the given function within a transaction
func (i *impl) WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := i.entConn.Tx(ctx)
	if err != nil {
		return err
	}

	// Execute the function
	if err := fn(tx); err != nil {
		// Roll back the transaction in case of error
		if rerr := tx.Rollback(); rerr != nil {
			// Return both errors if rollback fails
			return rerr
		}
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
