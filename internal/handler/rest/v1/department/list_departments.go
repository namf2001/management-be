package department

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DepartmentResponse represents a department in the system
// @name DepartmentResponse
type DepartmentResponse struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Engineering"`
	Description string    `json:"description" example:"Software Engineering Department"`
	CreatedAt   time.Time `json:"created_at" example:"2024-03-20T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-03-20T10:00:00Z"`
}

// ListDepartments
// @Summary      List all departments
// @Description  Get a list of all departments in the system
// @Tags         departments
// @Accept       json
// @Produce      json
// @Success      200  {object}  object{success=bool,data=object{departments=[]DepartmentResponse}}
// @Failure      500  {object}  object{success=bool,error=string}
// @Router       /api/departments [get]
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
