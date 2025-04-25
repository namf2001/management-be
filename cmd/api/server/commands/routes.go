package commands

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"management-be/internal/controller/department"
	"management-be/internal/controller/match"
	"management-be/internal/controller/player"
	"management-be/internal/controller/team"
	"management-be/internal/controller/user"
	v1 "management-be/internal/handler/rest/v1"
	"management-be/internal/pkg/middleware/auth"
	"management-be/internal/repository"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Middleware xử lý preflight OPTIONS
	r.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check & root
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	// Repositories & Controllers
	repoRegistry := repository.NewRegistry(s.db.Client())
	userController := user.NewController(repoRegistry)
	departmentController := department.NewController(repoRegistry)
	teamController := team.NewController(repoRegistry)
	playerController := player.NewController(repoRegistry)
	matchController := match.NewController(repoRegistry)
	handler := v1.NewHandler(userController, departmentController, teamController, playerController, matchController)

	// Routes with auth
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/change-password", auth.JWTAuthMiddleware(), handler.ChangePassword)
	}

	departmentGroup := r.Group("/api/departments", auth.JWTAuthMiddleware())
	{
		departmentGroup.GET("", handler.ListDepartments)
		departmentGroup.GET("/:id", handler.GetDepartment)
		departmentGroup.POST("", handler.CreateDepartment)
		departmentGroup.PUT("/:id", handler.UpdateDepartment)
		departmentGroup.DELETE("/:id", handler.DeleteDepartment)
	}

	teamGroup := r.Group("/api/teams", auth.JWTAuthMiddleware())
	{
		teamGroup.GET("", handler.ListTeams)
		teamGroup.GET("/:id", handler.GetTeam)
		teamGroup.POST("", handler.CreateTeam)
		teamGroup.PUT("/:id", handler.UpdateTeam)
		teamGroup.DELETE("/:id", handler.DeleteTeam)
		teamGroup.GET("/:id/statistics", handler.GetTeamStatistics)
	}

	matchGroup := r.Group("/api/matches", auth.JWTAuthMiddleware())
	{
		matchGroup.GET("", handler.ListMatches)
		matchGroup.GET("/:id", handler.GetMatch)
		matchGroup.POST("", handler.CreateMatch)
		//matchGroup.POST("/batch", handler.CreateManyMatches)
		matchGroup.PUT("/:id", handler.UpdateMatch)
		matchGroup.DELETE("/:id", handler.DeleteMatch)
		matchGroup.PUT("/:id/players", handler.UpdateMatchPlayers)
		matchGroup.GET("/:id/statistics", handler.GetMatchStatistics)
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
