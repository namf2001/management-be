package auth

import (
	"management-be/internal/controller/user"
)

type Handler struct {
	userCtrl user.Controller
}

func NewHandler(userCtrl user.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
	}
}
