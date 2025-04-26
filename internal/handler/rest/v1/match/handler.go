package match

import (
	"management-be/internal/controller/match"
	"management-be/internal/controller/player"
	"management-be/internal/controller/team"
)

type Handler struct {
	teamCtrl   team.Controller
	playerCtrl player.Controller
	matchCtrl  match.Controller
}

func NewHandler(teamCtrl team.Controller, playerCtrl player.Controller, matchCtrl match.Controller) Handler {
	return Handler{
		teamCtrl:   teamCtrl,
		playerCtrl: playerCtrl,
		matchCtrl:  matchCtrl,
	}
}
