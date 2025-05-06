package department

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteDepartment
// @Summary      Delete a department
// @Description  Delete a department by its ID
// @Tags         departments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Success      200  {object}  object{success=bool,message=string}
// @Failure      400  {object}  object{success=bool,error=string}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/departments/{id} [delete]
func (h Handler) DeleteDepartment(ctx *gin.Context) {
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

	// Delete department
	err = h.departmentCtrl.DeleteDepartment(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Department deleted successfully",
	})
}
