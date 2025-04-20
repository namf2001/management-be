package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

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
