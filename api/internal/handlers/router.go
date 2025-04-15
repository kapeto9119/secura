package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/secura/api/internal/config"
	"github.com/secura/api/internal/middlewares"
)

// SetupRouter configures the Gin router and registers all routes
func SetupRouter(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new router
	router := gin.New()

	// Register global middlewares
	router.Use(gin.Recovery())

	// Define routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		public := v1.Group("/")
		{
			public.POST("/auth/login", Login(cfg))
			public.GET("/health", HealthCheck(cfg))
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middlewares.JWTAuth(cfg))
		{
			// User routes
			protected.GET("/user", GetUser())

			// LLM routes
			llmRoutes := protected.Group("/llm")
			{
				llmRoutes.POST("/completion", LLMCompletion(cfg, logger))
				llmRoutes.POST("/chat", LLMChat(cfg, logger))
			}

			// Audit routes
			auditRoutes := protected.Group("/audit")
			{
				auditRoutes.GET("/logs", GetAuditLogs())
				auditRoutes.GET("/logs/:id", GetAuditLog())
			}
		}
	}

	return router
}
