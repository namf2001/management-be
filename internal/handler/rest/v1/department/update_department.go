package department

import (
	"management-be/internal/pkg/unit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateDepartmentRequest represents the request body for updating a department
// @name UpdateDepartmentRequest
type UpdateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100" example:"Engineering"`
	Description string `json:"description" validate:"required" example:"Software Engineering Department"`
}

// UpdateDepartment
// @Summary      Update a department
// @Description  Update an existing department's information
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Param        department  body      UpdateDepartmentRequest  true  "Department information"
// @Success      200  {object}  object{success=bool,data=Department}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/departments/{id} [put]
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
	if !unit.ValidateJSON(ctx, &req) {
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
		"data": Department{
			ID:          department.ID,
			Name:        department.Name,
			Description: department.Description,
			CreatedAt:   department.CreatedAt,
			UpdatedAt:   department.UpdatedAt,
		},
	})
}
