package department

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateDepartment handles the request to create a new department
func (h Handler) CreateDepartment(ctx *gin.Context) {
	var req CreateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Create department
	department, err := h.departmentCtrl.CreateDepartment(ctx.Request.Context(), req.Name, req.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create department",
		})
		return
	}

	// Return created department
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": DepartmentResponse{
			ID:          department.ID,
			Name:        department.Name,
			Description: department.Description,
			CreatedAt:   department.CreatedAt,
			UpdatedAt:   department.UpdatedAt,
		},
	})
}
