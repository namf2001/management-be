package auth

import (
	"management-be/internal/pkg/unit"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents the request body for user login
// @name LoginRequest
type LoginRequest struct {
	Username string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// LoginResponse represents the response body for user login
// @name LoginResponse
type LoginResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}

// @Summary      Login user
// @Description  Login user with username and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "User login credentials"
// @Success      200         {object}  LoginResponse
// @Failure      400         {object}  object{error=string}
// @Failure      401         {object}  object{error=string}
// @Failure      500         {object}  object{error=string}
// @Router       /api/users/login [post]
func (h Handler) Login(ctx *gin.Context) {
	var req LoginRequest

	// Use the validator package
	if !unit.ValidateJSON(ctx, &req) {
		return
	}

	token, user, err := h.userCtrl.Login(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	response := LoginResponse{
		Token: token,
		User: UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
