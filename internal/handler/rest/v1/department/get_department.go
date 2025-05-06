package department

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDepartment
// @Summary      Get a department by ID
// @Description  Get detailed information about a specific department
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Success      200  {object}  object{success=bool,data=DepartmentResponse}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      404  {object}  object{success=bool,error=string}
// @Router       /api/departments/{id} [get]
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
