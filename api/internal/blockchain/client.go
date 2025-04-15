package blockchain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// Client represents a client for interacting with the blockchain
type Client struct {
	ethClient    *ethclient.Client
	contractAddr common.Address
	logger       *zap.Logger
}

// AuditLogData represents the data for an audit log
type AuditLogData struct {
	UserID       string                 `json:"user_id"`
	ActionType   string                 `json:"action_type"`
	RequestData  map[string]interface{} `json:"request_data"`
	ResponseData map[string]interface{} `json:"response_data"`
	Metadata     map[string]interface{} `json:"metadata"`
}

// NewClient creates a new blockchain client
func NewClient(nodeURL string, contractAddress string, logger *zap.Logger) (*Client, error) {
	// Connect to the Ethereum node
	ethClient, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	// Validate contract address
	if !common.IsHexAddress(contractAddress) {
		return nil, fmt.Errorf("invalid contract address: %s", contractAddress)
	}
	addr := common.HexToAddress(contractAddress)

	client := &Client{
		ethClient:    ethClient,
		contractAddr: addr,
		logger:       logger,
	}

	return client, nil
}

// RecordAuditLog records an audit log to the blockchain
func (c *Client) RecordAuditLog(ctx context.Context, data AuditLogData) (string, error) {
	// Create a JSON representation of the request and response data
	requestJSON, err := json.Marshal(data.RequestData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	responseJSON, err := json.Marshal(data.ResponseData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response data: %w", err)
	}

	// Create content hash
	combinedData := append(requestJSON, responseJSON...)
	contentHash := sha256.Sum256(combinedData)
	contentHashHex := hex.EncodeToString(contentHash[:])

	// Create metadata
	metadata, err := json.Marshal(data.Metadata)
	if err != nil {
		return "", fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// For MVP, just log the data and return a mock transaction hash
	// TODO: Implement actual blockchain integration
	c.logger.Info("Recording audit log to blockchain",
		zap.String("user_id", data.UserID),
		zap.String("action_type", data.ActionType),
		zap.String("content_hash", contentHashHex),
		zap.String("metadata", string(metadata)),
	)

	// Return a mock transaction hash for now
	return fmt.Sprintf("0x%s", hex.EncodeToString(contentHash[:])[:40]), nil
}

// VerifyContentHash verifies if a content hash exists in the audit trail
func (c *Client) VerifyContentHash(ctx context.Context, contentHash string) (bool, error) {
	// For MVP, just log the verification attempt and return a mock result
	// TODO: Implement actual blockchain verification
	c.logger.Info("Verifying content hash in blockchain",
		zap.String("content_hash", contentHash),
	)

	// Always return true for MVP
	return true, nil
}

// GetAuditLogsByUser retrieves all audit logs for a specific user
func (c *Client) GetAuditLogsByUser(ctx context.Context, userID string) ([]map[string]interface{}, error) {
	// For MVP, just log the request and return mock data
	// TODO: Implement actual blockchain queries
	c.logger.Info("Getting audit logs for user",
		zap.String("user_id", userID),
	)

	// Return mock data
	return []map[string]interface{}{
		{
			"index":     0,
			"user_id":   userID,
			"action":    "completion",
			"hash":      "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			"timestamp": 1633042800,
			"metadata":  `{"model":"gpt-4","token_count":150}`,
		},
		{
			"index":     1,
			"user_id":   userID,
			"action":    "chat",
			"hash":      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
			"timestamp": 1633046400,
			"metadata":  `{"model":"gpt-3.5-turbo","token_count":200}`,
		},
	}, nil
}

// GenerateContentHash generates a content hash from request and response data
func GenerateContentHash(requestData, responseData map[string]interface{}) (string, error) {
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response data: %w", err)
	}

	combinedData := append(requestJSON, responseJSON...)
	hash := sha256.Sum256(combinedData)
	return hex.EncodeToString(hash[:]), nil
}
