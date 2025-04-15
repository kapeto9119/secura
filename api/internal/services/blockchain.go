package services

import (
	"context"
	"fmt"

	"github.com/secura/api/internal/blockchain"
	"go.uber.org/zap"
)

// BlockchainService handles interaction with the blockchain
type BlockchainService struct {
	client *blockchain.Client
	logger *zap.Logger
}

// NewBlockchainService creates a new blockchain service
func NewBlockchainService(nodeURL string, contractAddress string, logger *zap.Logger) (*BlockchainService, error) {
	client, err := blockchain.NewClient(nodeURL, contractAddress, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain client: %w", err)
	}

	return &BlockchainService{
		client: client,
		logger: logger,
	}, nil
}

// RecordLLMInteraction records an LLM interaction (request and response) to the blockchain
func (s *BlockchainService) RecordLLMInteraction(
	ctx context.Context,
	userID string,
	actionType string,
	requestData map[string]interface{},
	responseData map[string]interface{},
	metadata map[string]interface{},
) (string, error) {
	// Create audit log data
	logData := blockchain.AuditLogData{
		UserID:       userID,
		ActionType:   actionType,
		RequestData:  requestData,
		ResponseData: responseData,
		Metadata:     metadata,
	}

	// Record to blockchain
	txHash, err := s.client.RecordAuditLog(ctx, logData)
	if err != nil {
		s.logger.Error("Failed to record audit log",
			zap.Error(err),
			zap.String("user_id", userID),
			zap.String("action_type", actionType),
		)
		return "", fmt.Errorf("failed to record audit log: %w", err)
	}

	return txHash, nil
}

// VerifyContentHash verifies if a content hash exists in the audit trail
func (s *BlockchainService) VerifyContentHash(ctx context.Context, contentHash string) (bool, error) {
	exists, err := s.client.VerifyContentHash(ctx, contentHash)
	if err != nil {
		s.logger.Error("Failed to verify content hash",
			zap.Error(err),
			zap.String("content_hash", contentHash),
		)
		return false, fmt.Errorf("failed to verify content hash: %w", err)
	}

	return exists, nil
}

// GetUserAuditLogs retrieves all audit logs for a specific user
func (s *BlockchainService) GetUserAuditLogs(ctx context.Context, userID string) ([]map[string]interface{}, error) {
	logs, err := s.client.GetAuditLogsByUser(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to get user audit logs",
			zap.Error(err),
			zap.String("user_id", userID),
		)
		return nil, fmt.Errorf("failed to get user audit logs: %w", err)
	}

	return logs, nil
}
