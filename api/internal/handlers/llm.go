package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/secura/api/internal/config"
	"github.com/secura/api/internal/services"
)

// CompletionRequest represents a request to the completion endpoint
type CompletionRequest struct {
	Prompt      string  `json:"prompt" binding:"required"`
	Model       string  `json:"model" binding:"required"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

// ChatRequest represents a request to the chat endpoint
type ChatRequest struct {
	Messages    []Message `json:"messages" binding:"required"`
	Model       string    `json:"model" binding:"required"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// LLMCompletion handles completion requests
func LLMCompletion(cfg *config.Config, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CompletionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Get user ID from context
		userID, _ := c.Get("userID")
		logger.Info("Processing completion request", zap.String("user_id", userID.(string)), zap.String("model", req.Model))

		// Anonymize the prompt
		anonService := services.NewAnonymizationService(cfg.NLPServiceURL)
		anonymizedPrompt, err := anonService.AnonymizeText(req.Prompt)
		if err != nil {
			logger.Error("Failed to anonymize prompt", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to process request",
			})
			return
		}

		// Forward to OpenAI
		openaiReq := map[string]interface{}{
			"prompt":      anonymizedPrompt,
			"model":       req.Model,
			"max_tokens":  req.MaxTokens,
			"temperature": req.Temperature,
		}

		resp, err := forwardToOpenAI(cfg.OpenAIAPIKey, "https://api.openai.com/v1/completions", openaiReq)
		if err != nil {
			logger.Error("Failed to call OpenAI", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to process request",
			})
			return
		}

		// TODO: Store in blockchain audit trail

		// Return response to client
		c.JSON(http.StatusOK, resp)
	}
}

// LLMChat handles chat requests
func LLMChat(cfg *config.Config, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Get user ID from context
		userID, _ := c.Get("userID")
		logger.Info("Processing chat request", zap.String("user_id", userID.(string)), zap.String("model", req.Model))

		// Anonymize the messages
		anonService := services.NewAnonymizationService(cfg.NLPServiceURL)
		for i, msg := range req.Messages {
			anonymizedContent, err := anonService.AnonymizeText(msg.Content)
			if err != nil {
				logger.Error("Failed to anonymize message", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to process request",
				})
				return
			}
			req.Messages[i].Content = anonymizedContent
		}

		// Forward to OpenAI
		openaiReq := map[string]interface{}{
			"messages":    req.Messages,
			"model":       req.Model,
			"max_tokens":  req.MaxTokens,
			"temperature": req.Temperature,
		}

		resp, err := forwardToOpenAI(cfg.OpenAIAPIKey, "https://api.openai.com/v1/chat/completions", openaiReq)
		if err != nil {
			logger.Error("Failed to call OpenAI", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to process request",
			})
			return
		}

		// TODO: Store in blockchain audit trail

		// Return response to client
		c.JSON(http.StatusOK, resp)
	}
}

// Helper function to forward requests to OpenAI
func forwardToOpenAI(apiKey string, url string, reqBody map[string]interface{}) (map[string]interface{}, error) {
	// Marshal request body
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
