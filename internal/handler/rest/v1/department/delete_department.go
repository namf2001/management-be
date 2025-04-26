package department

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteDepartment handles the request to delete a department by ID
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
