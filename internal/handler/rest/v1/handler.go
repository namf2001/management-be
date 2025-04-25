package v1

import (
	"management-be/internal/controller/department"
	"management-be/internal/controller/match"
	"management-be/internal/controller/player"
	"management-be/internal/controller/team"
	"management-be/internal/controller/user"
)

type Handler struct {
	userCtrl       user.Controller
	departmentCtrl department.Controller
	teamCtrl       team.Controller
	playerCtrl     player.Controller
	matchCtrl      match.Controller
}

func NewHandler(userCtrl user.Controller, departmentCtrl department.Controller, teamCtrl team.Controller, playerCtrl player.Controller, matchCtrl match.Controller) Handler {
	return Handler{
		userCtrl:       userCtrl,
		departmentCtrl: departmentCtrl,
		teamCtrl:       teamCtrl,
		playerCtrl:     playerCtrl,
		matchCtrl:      matchCtrl,
	}
}
