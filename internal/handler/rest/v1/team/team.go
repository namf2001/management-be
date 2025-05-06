package team

import (
	"management-be/internal/controller/team"
)

// UpdateHandler to include team controller
func (h Handler) UpdateTeamController(teamCtrl team.Controller) {
	h.teamCtrl = teamCtrl
}
