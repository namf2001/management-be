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
	userHandler "management-be/internal/handler/rest/v1/auth"
	departmentHandler "management-be/internal/handler/rest/v1/department"
	matchHandler "management-be/internal/handler/rest/v1/match"
	teamHandler "management-be/internal/handler/rest/v1/team"
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

	// Initialize handlers
	userH := userHandler.NewHandler(userController)
	teamH := teamHandler.NewHandler(teamController, playerController, matchController)
	departmentH := departmentHandler.NewHandler(departmentController)
	matchH := matchHandler.NewHandler(teamController, playerController, matchController)

	// Routes with auth
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", userH.Register)
		userGroup.POST("/login", userH.Login)
		userGroup.POST("/change-password", auth.JWTAuthMiddleware(), userH.ChangePassword)
	}

	departmentGroup := r.Group("/api/departments")
	{
		departmentGroup.GET("", departmentH.ListDepartments)
		departmentGroup.GET("/:id", departmentH.GetDepartment)
		departmentGroup.POST("", departmentH.CreateDepartment)
		departmentGroup.PUT("/:id", departmentH.UpdateDepartment)
		departmentGroup.DELETE("/:id", departmentH.DeleteDepartment)
	}

	teamGroup := r.Group("/api/teams")
	{
		teamGroup.GET("", teamH.ListTeams)
		teamGroup.GET("/:id", teamH.GetTeam)
		teamGroup.POST("", teamH.CreateTeam)
		teamGroup.PUT("/:id", teamH.UpdateTeam)
		teamGroup.DELETE("/:id", teamH.DeleteTeam)
		teamGroup.GET("/:id/statistics", teamH.GetTeamStatistics)
	}

	matchGroup := r.Group("/api/matches")
	{
		matchGroup.GET("", matchH.ListMatches)
		matchGroup.GET("/:id", matchH.GetMatch)
		matchGroup.POST("", matchH.CreateMatch)
		//matchGroup.POST("/batch", handler.CreateManyMatches)
		matchGroup.PUT("/:id", matchH.UpdateMatch)
		matchGroup.DELETE("/:id", matchH.DeleteMatch)
		matchGroup.PUT("/:id/players", matchH.UpdateMatchPlayers)
		matchGroup.GET("/:id/statistics", matchH.GetMatchStatistics)
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
