// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "match_date", Type: field.TypeTime},
		{Name: "venue", Type: field.TypeString, Nullable: true},
		{Name: "is_home_game", Type: field.TypeBool},
		{Name: "our_score", Type: field.TypeInt32, Nullable: true},
		{Name: "opponent_score", Type: field.TypeInt32, Nullable: true},
		{Name: "status", Type: field.TypeString},
		{Name: "notes", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "opponent_team_id", Type: field.TypeInt, Nullable: true},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_teams_matches",
				Columns:    []*schema.Column{MatchesColumns[11]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MatchPlayersColumns holds the columns for the "match_players" table.
	MatchPlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "minutes_played", Type: field.TypeInt32},
		{Name: "goals_scored", Type: field.TypeInt32},
		{Name: "assists", Type: field.TypeInt32},
		{Name: "yellow_cards", Type: field.TypeInt32},
		{Name: "red_card", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "match_id", Type: field.TypeInt, Nullable: true},
		{Name: "player_id", Type: field.TypeInt, Nullable: true},
	}
	// MatchPlayersTable holds the schema information for the "match_players" table.
	MatchPlayersTable = &schema.Table{
		Name:       "match_players",
		Columns:    MatchPlayersColumns,
		PrimaryKey: []*schema.Column{MatchPlayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "match_players_matches_match_players",
				Columns:    []*schema.Column{MatchPlayersColumns[9]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "match_players_players_match_players",
				Columns:    []*schema.Column{MatchPlayersColumns[10]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MatchesGatewayColumns holds the columns for the "matches_gateway" table.
	MatchesGatewayColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "competition_name", Type: field.TypeString},
		{Name: "season_start_date", Type: field.TypeTime},
		{Name: "match_date", Type: field.TypeTime},
		{Name: "home_team_name", Type: field.TypeString},
		{Name: "home_team_short_name", Type: field.TypeString},
		{Name: "home_team_logo", Type: field.TypeString},
		{Name: "away_team_name", Type: field.TypeString},
		{Name: "away_team_short_name", Type: field.TypeString},
		{Name: "away_team_logo", Type: field.TypeString},
		{Name: "home_score", Type: field.TypeInt32, Nullable: true},
		{Name: "away_score", Type: field.TypeInt32, Nullable: true},
		{Name: "status", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// MatchesGatewayTable holds the schema information for the "matches_gateway" table.
	MatchesGatewayTable = &schema.Table{
		Name:       "matches_gateway",
		Columns:    MatchesGatewayColumns,
		PrimaryKey: []*schema.Column{MatchesGatewayColumns[0]},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "full_name", Type: field.TypeString, Unique: true},
		{Name: "jersey_number", Type: field.TypeInt32, Unique: true, Nullable: true},
		{Name: "position", Type: field.TypeString},
		{Name: "date_of_birth", Type: field.TypeTime, Nullable: true},
		{Name: "height_cm", Type: field.TypeInt32, Nullable: true},
		{Name: "weight_kg", Type: field.TypeInt32, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "players_departments_players",
				Columns:    []*schema.Column{PlayersColumns[13]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PlayerStatisticsColumns holds the columns for the "player_statistics" table.
	PlayerStatisticsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "total_matches", Type: field.TypeInt32},
		{Name: "total_minutes_played", Type: field.TypeInt32},
		{Name: "total_goals", Type: field.TypeInt32},
		{Name: "total_assists", Type: field.TypeInt32},
		{Name: "total_yellow_cards", Type: field.TypeInt32},
		{Name: "total_red_cards", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "player_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PlayerStatisticsTable holds the schema information for the "player_statistics" table.
	PlayerStatisticsTable = &schema.Table{
		Name:       "player_statistics",
		Columns:    PlayerStatisticsColumns,
		PrimaryKey: []*schema.Column{PlayerStatisticsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "player_statistics_players_player_statistic",
				Columns:    []*schema.Column{PlayerStatisticsColumns[10]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SchemaMigrationsColumns holds the columns for the "schema_migrations" table.
	SchemaMigrationsColumns = []*schema.Column{
		{Name: "version", Type: field.TypeInt, Increment: true},
		{Name: "dirty", Type: field.TypeBool},
	}
	// SchemaMigrationsTable holds the schema information for the "schema_migrations" table.
	SchemaMigrationsTable = &schema.Table{
		Name:       "schema_migrations",
		Columns:    SchemaMigrationsColumns,
		PrimaryKey: []*schema.Column{SchemaMigrationsColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "company_name", Type: field.TypeString, Nullable: true},
		{Name: "contact_person", Type: field.TypeString, Nullable: true},
		{Name: "contact_phone", Type: field.TypeString, Nullable: true},
		{Name: "contact_email", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// TeamFeesColumns holds the columns for the "team_fees" table.
	TeamFeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "payment_date", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// TeamFeesTable holds the schema information for the "team_fees" table.
	TeamFeesTable = &schema.Table{
		Name:       "team_fees",
		Columns:    TeamFeesColumns,
		PrimaryKey: []*schema.Column{TeamFeesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "full_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
		MatchesTable,
		MatchPlayersTable,
		MatchesGatewayTable,
		PlayersTable,
		PlayerStatisticsTable,
		SchemaMigrationsTable,
		TeamsTable,
		TeamFeesTable,
		UsersTable,
	}
)

func init() {
	MatchesTable.ForeignKeys[0].RefTable = TeamsTable
	MatchPlayersTable.ForeignKeys[0].RefTable = MatchesTable
	MatchPlayersTable.ForeignKeys[1].RefTable = PlayersTable
	MatchesGatewayTable.Annotation = &entsql.Annotation{
		Table: "matches_gateway",
	}
	PlayersTable.ForeignKeys[0].RefTable = DepartmentsTable
	PlayerStatisticsTable.ForeignKeys[0].RefTable = PlayersTable
}
