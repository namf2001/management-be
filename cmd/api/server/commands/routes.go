package commands

import (
	v1 "management-be/internal/handler/rest/v1"
	"management-be/internal/pkg/middleware/auth"
	"net/http"

	"management-be/internal/controller/department"
	"management-be/internal/controller/team"
	"management-be/internal/controller/user"
	"management-be/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	// Initialize repositories and controllers
	repoRegistry := repository.NewRegistry(s.db.Client())
	userController := user.NewController(repoRegistry)
	departmentController := department.NewController(repoRegistry)
	teamController := team.NewController(repoRegistry)
	handler := v1.NewHandler(userController, departmentController, teamController)

	// User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/change-password", auth.JWTAuthMiddleware(), handler.ChangePassword)
	}

	// Department routes
	departmentGroup := r.Group("/api/departments")
	{
		departmentGroup.GET("", auth.JWTAuthMiddleware(), handler.ListDepartments)
		departmentGroup.GET("/:id", auth.JWTAuthMiddleware(), handler.GetDepartment)
		departmentGroup.POST("", auth.JWTAuthMiddleware(), handler.CreateDepartment)
		departmentGroup.PUT("/:id", auth.JWTAuthMiddleware(), handler.UpdateDepartment)
		departmentGroup.DELETE("/:id", auth.JWTAuthMiddleware(), handler.DeleteDepartment)
	}

	// Team routes
	teamGroup := r.Group("/api/teams")
	{
		teamGroup.GET("", auth.JWTAuthMiddleware(), handler.ListTeams)
		teamGroup.GET("/:id", auth.JWTAuthMiddleware(), handler.GetTeam)
		teamGroup.POST("", auth.JWTAuthMiddleware(), handler.CreateTeam)
		teamGroup.PUT("/:id", auth.JWTAuthMiddleware(), handler.UpdateTeam)
		teamGroup.DELETE("/:id", auth.JWTAuthMiddleware(), handler.DeleteTeam)
		teamGroup.GET("/:id/statistics", auth.JWTAuthMiddleware(), handler.GetTeamStatistics)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
