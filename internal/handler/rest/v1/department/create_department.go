package department

import (
	"github.com/gin-gonic/gin"
	v1 "management-be/internal/handler/rest/v1"
	"net/http"
	"time"
)

// CreateDepartmentRequest represents the request body for creating a department
type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateDepartmentResponse represents the response format for a department
type CreateDepartmentResponse struct {
	Success bool          `json:"success"`
	Data    *Department   `json:"data,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   *v1.ErrorInfo `json:"error,omitempty"`
}

type Department struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateDepartment handles the request to create a new department
func (h Handler) CreateDepartment(ctx *gin.Context) {
	var req CreateDepartmentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateDepartmentResponse{
			Success: false,
			Message: "Invalid request payload",
			Error: &v1.ErrorInfo{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
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
