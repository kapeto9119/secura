package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/secura/api/internal/config"
	"github.com/secura/api/internal/services"
)

// MockAuditLog represents a mock audit log entry
type MockAuditLog struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Action        string    `json:"action"`
	Timestamp     time.Time `json:"timestamp"`
	BlockchainTxn string    `json:"blockchain_txn"`
	Status        string    `json:"status"`
	Details       gin.H     `json:"details"`
}

// GetAuditLogs returns a handler for retrieving audit logs
func GetAuditLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// Get blockchain service from the context (set in router setup)
		blockchainService, exists := c.Get("blockchainService")
		if !exists {
			// Fall back to mock data if service not available
			logs := getMockAuditLogs(userID.(string))
			c.JSON(http.StatusOK, gin.H{
				"logs":  logs,
				"total": len(logs),
			})
			return
		}

		// Get logs from blockchain
		ctx := context.Background()
		logs, err := blockchainService.(*services.BlockchainService).GetUserAuditLogs(ctx, userID.(string))
		if err != nil {
			// Log error and fall back to mock data
			logger, _ := c.Get("logger")
			if logger != nil {
				logger.(*zap.Logger).Error("Failed to get audit logs from blockchain", zap.Error(err))
			}

			logs := getMockAuditLogs(userID.(string))
			c.JSON(http.StatusOK, gin.H{
				"logs":  logs,
				"total": len(logs),
				"note":  "Using mock data due to blockchain service error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"logs":  logs,
			"total": len(logs),
		})
	}
}

// GetAuditLog returns a handler for retrieving a specific audit log
func GetAuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := c.Param("id")

		// Get user ID from context
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// For MVP, just return mock data
		// In a real implementation, we would query the blockchain
		if logID != "log-001" && logID != "log-002" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Audit log not found",
			})
			return
		}

		var log MockAuditLog
		if logID == "log-001" {
			log = MockAuditLog{
				ID:            "log-001",
				UserID:        userID.(string),
				Action:        "completion",
				Timestamp:     time.Now().Add(-1 * time.Hour),
				BlockchainTxn: "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
				Status:        "success",
				Details: gin.H{
					"model":           "gpt-4",
					"prompt_hash":     "sha256:1234567890abcdef",
					"response_hash":   "sha256:abcdef1234567890",
					"token_count":     150,
					"anonymized_data": true,
				},
			}
		} else {
			log = MockAuditLog{
				ID:            "log-002",
				UserID:        userID.(string),
				Action:        "chat",
				Timestamp:     time.Now().Add(-2 * time.Hour),
				BlockchainTxn: "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				Status:        "success",
				Details: gin.H{
					"model":           "gpt-3.5-turbo",
					"message_hash":    "sha256:fedcba0987654321",
					"response_hash":   "sha256:1234567890fedcba",
					"token_count":     200,
					"anonymized_data": true,
				},
			}
		}

		c.JSON(http.StatusOK, log)
	}
}

// Setup audit handlers and blockchain service
func SetupAuditHandlers(router *gin.RouterGroup, cfg *config.Config, logger *zap.Logger) {
	// Initialize blockchain service if enabled
	if cfg.BlockchainNodeURL != "" {
		blockchainService, err := services.NewBlockchainService(
			cfg.BlockchainNodeURL,
			cfg.BlockchainContractAddress,
			logger,
		)

		if err != nil {
			logger.Error("Failed to initialize blockchain service for audit handlers", zap.Error(err))
		} else {
			// Add blockchain service to the router context
			router.Use(func(c *gin.Context) {
				c.Set("blockchainService", blockchainService)
				c.Set("logger", logger)
				c.Next()
			})
		}
	}

	// Set up routes
	router.GET("/logs", GetAuditLogs())
	router.GET("/logs/:id", GetAuditLog())
}

// Helper function to get mock audit logs
func getMockAuditLogs(userID string) []MockAuditLog {
	return []MockAuditLog{
		{
			ID:            "log-001",
			UserID:        userID,
			Action:        "completion",
			Timestamp:     time.Now().Add(-1 * time.Hour),
			BlockchainTxn: "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			Status:        "success",
			Details: gin.H{
				"model":           "gpt-4",
				"prompt_hash":     "sha256:1234567890abcdef",
				"response_hash":   "sha256:abcdef1234567890",
				"token_count":     150,
				"anonymized_data": true,
			},
		},
		{
			ID:            "log-002",
			UserID:        userID,
			Action:        "chat",
			Timestamp:     time.Now().Add(-2 * time.Hour),
			BlockchainTxn: "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
			Status:        "success",
			Details: gin.H{
				"model":           "gpt-3.5-turbo",
				"message_hash":    "sha256:fedcba0987654321",
				"response_hash":   "sha256:1234567890fedcba",
				"token_count":     200,
				"anonymized_data": true,
			},
		},
	}
}
