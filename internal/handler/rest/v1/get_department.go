package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetDepartment handles the request to get a department by ID
func (h Handler) GetDepartment(ctx *gin.Context) {
	// Get department ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid department ID",
		})
		return
	}

	// Get department from controller
	department, err := h.departmentCtrl.GetDepartmentByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Department not found",
		})
		return
	}

	// Return department
	ctx.JSON(http.StatusOK, gin.H{
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
