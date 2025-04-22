package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRequest represents the request body for user registration
// @name RegisterRequest
type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"john_doe"`
	Password string `json:"password" binding:"required" example:"password123"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	FullName string `json:"full_name" example:"John Doe"`
}

// RegisterResponse represents the response body for user registration
// @name RegisterResponse
type RegisterResponse struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"john@example.com"`
	FullName string `json:"full_name" example:"John Doe"`
}

// Register
// @Summary      Register a new user
// @Description  Register a new user with username, password, email and full name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      RegisterRequest  true  "User registration information"
// @Success      201  {object}  RegisterResponse
// @Failure      400  {object}  object{error=string}
// @Failure      500  {object}  object{error=string}
// @Router       /api/users/register [post]
func (h Handler) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userCtrl.CreateUser(ctx.Request.Context(), req.Username, req.Email, req.Password, req.FullName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userFound, err := h.userCtrl.GetUserByID(ctx.Request.Context(), user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user details"})
		return
	}

	response := RegisterResponse{
		ID:       userFound.ID,
		Username: userFound.Username,
		Email:    userFound.Email,
		FullName: userFound.FullName,
	}

	ctx.JSON(http.StatusCreated, response)
}
