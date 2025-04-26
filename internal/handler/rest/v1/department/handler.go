package department

import (
	"management-be/internal/controller/department"
)

type Handler struct {
	departmentCtrl department.Controller
}

func NewHandler(departmentCtrl department.Controller) Handler {
	return Handler{
		departmentCtrl: departmentCtrl,
	}
}
