package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DepartmentResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListDepartments handles the request to list all departments
func (h Handler) ListDepartments(ctx *gin.Context) {
	departments, err := h.departmentCtrl.GetAllDepartments(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve departments",
		})
		return
	}

	// Convert departments to response format
	response := make([]DepartmentResponse, len(departments))
	for i, dept := range departments {
		response[i] = DepartmentResponse{
			ID:          dept.ID,
			Name:        dept.Name,
			Description: dept.Description,
			CreatedAt:   dept.CreatedAt,
			UpdatedAt:   dept.UpdatedAt,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"departments": response,
		},
	})
}
