package team_fee

import (
	"management-be/internal/controller/team_fee"
)

// Handler represents the team fee handler
type Handler struct {
	teamFeeCtrl team_fee.Controller
}

// NewHandler creates a new team fee handler
func NewHandler(teamFeeCtrl team_fee.Controller) Handler {
	return Handler{
		teamFeeCtrl: teamFeeCtrl,
	}
}
