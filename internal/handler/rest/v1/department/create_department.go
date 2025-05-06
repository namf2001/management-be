package department

import (
	v1 "management-be/internal/handler/rest/v1"
	"management-be/internal/pkg/unit"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateDepartmentRequest represents the request body for creating a department
// @name CreateDepartmentRequest
type CreateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100" example:"Engineering"`
	Description string `json:"description" validate:"required" example:"Software Engineering Department"`
}

// CreateDepartmentResponse represents the response format for a department
// @name CreateDepartmentResponse
type CreateDepartmentResponse struct {
	Success bool          `json:"success"`
	Data    *Department   `json:"data,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   *v1.ErrorInfo `json:"error,omitempty"`
}

// Department represents a department in the system
// @name Department
type Department struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Engineering"`
	Description string    `json:"description" example:"Software Engineering Department"`
	CreatedAt   time.Time `json:"created_at" example:"2024-03-20T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-03-20T10:00:00Z"`
}

// CreateDepartment
// @Summary      Create a new department
// @Description  Create a new department with name and description
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        department  body      CreateDepartmentRequest  true  "Department information"
// @Success      201  {object}  CreateDepartmentResponse
// @Failure      400  {object}  CreateDepartmentResponse
// @Failure      500  {object}  CreateDepartmentResponse
// @Router       /api/departments [post]
func (h Handler) CreateDepartment(ctx *gin.Context) {
	var req CreateDepartmentRequest

	// Use the validator package
	if !unit.ValidateJSON(ctx, &req) {
		return
	}

	// Create department
	department, err := h.departmentCtrl.CreateDepartment(ctx.Request.Context(), req.Name, req.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateDepartmentResponse{
			Success: false,
			Message: "Failed to create department",
			Error: &v1.ErrorInfo{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		})
		return
	}

	// Return created department
	ctx.JSON(http.StatusCreated, CreateDepartmentResponse{
		Success: true,
		Data: &Department{
			ID:          department.ID,
			Name:        department.Name,
			Description: department.Description,
			CreatedAt:   department.CreatedAt,
			UpdatedAt:   department.UpdatedAt,
		},
		Message: "Department created successfully",
	})
}
