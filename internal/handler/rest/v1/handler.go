package v1

import (
	"management-be/internal/controller/department"
	"management-be/internal/controller/user"
)

type Handler struct {
	userCtrl       user.Controller
	departmentCtrl department.Controller
}

func NewHandler(userCtrl user.Controller, departmentCtrl department.Controller) Handler {
	return Handler{
		userCtrl:       userCtrl,
		departmentCtrl: departmentCtrl,
	}
}
