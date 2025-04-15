package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/secura/api/internal/models"
)

// GetUser returns a handler for retrieving the current user
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// TODO: Fetch user data from database
		// For now, just return a mock user
		user := models.User{
			ID:       userID.(string),
			Username: "admin",
			Email:    "admin@example.com",
			Role:     "admin",
		}

		c.JSON(http.StatusOK, user)
	}
}
