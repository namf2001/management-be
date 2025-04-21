package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	UserIdKey = "user_id"
)

// ChangePasswordRequest represents the request body for changing a password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

// ChangePassword handles the request to change a user's password.
func (h Handler) ChangePassword(ctx *gin.Context) {
	// Get user ID from the JWT token
	userIDStr, exists := ctx.Get(UserIdKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
		})
		return
	}

	userID, ok := userIDStr.(int)
	if !ok {
		// Try to convert from string if it's not already an int
		if userIDStr, ok := userIDStr.(string); ok {
			var err error
			userID, err = strconv.Atoi(userIDStr)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Invalid user ID",
				})
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Invalid user ID",
			})
			return
		}
	}

	var req ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.userCtrl.UpdatePassword(ctx.Request.Context(), userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to change password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password changed successfully",
	})
}
