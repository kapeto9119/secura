package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/secura/api/internal/config"
)

// HealthCheck returns a handler for the health endpoint
func HealthCheck(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"version":     "0.1.0",
			"environment": cfg.Environment,
			"timestamp":   time.Now().Format(time.RFC3339),
		})
	}
}
