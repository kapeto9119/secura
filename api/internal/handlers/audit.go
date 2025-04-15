package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
		// TODO: Implement actual blockchain querying
		// For now, return mock data

		logs := []MockAuditLog{
			{
				ID:            "log-001",
				UserID:        "user-123",
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
				UserID:        "user-123",
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

		c.JSON(http.StatusOK, gin.H{
			"logs":  logs,
			"total": 2,
		})
	}
}

// GetAuditLog returns a handler for retrieving a specific audit log
func GetAuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := c.Param("id")

		// TODO: Implement actual blockchain querying
		// For now, return mock data

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
				UserID:        "user-123",
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
				UserID:        "user-123",
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
