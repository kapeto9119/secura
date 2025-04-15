package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/secura/api/internal/config"
	"github.com/secura/api/internal/models"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login returns a handler for the login endpoint
func Login(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// TODO: Replace with actual authentication logic
		if req.Username != "admin" || req.Password != "password" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}

		// Create a new token
		now := time.Now()
		expirationTime := now.Add(time.Duration(cfg.JWTExpiryHours) * time.Hour)

		claims := jwt.MapClaims{
			"sub":  "user-123", // User ID
			"name": req.Username,
			"role": "admin",
			"iat":  now.Unix(),
			"exp":  expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}

		// Return the token
		c.JSON(http.StatusOK, gin.H{
			"token":      tokenString,
			"expires_at": expirationTime.Format(time.RFC3339),
			"user": models.User{
				ID:       "user-123",
				Username: req.Username,
				Role:     "admin",
			},
		})
	}
}
