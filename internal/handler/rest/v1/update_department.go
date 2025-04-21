package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// UpdateDepartment handles the request to update an existing department
func (h Handler) UpdateDepartment(ctx *gin.Context) {
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

	var req UpdateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update department
	department, err := h.departmentCtrl.UpdateDepartment(ctx.Request.Context(), id, req.Name, req.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update department",
		})
		return
	}

	// Return updated department
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
