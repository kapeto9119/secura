package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// AnonymizationService handles anonymization of sensitive data
type AnonymizationService struct {
	baseURL string
}

// AnonymizeRequest represents a request to the anonymization service
type AnonymizeRequest struct {
	Text string `json:"text"`
}

// AnonymizeResponse represents a response from the anonymization service
type AnonymizeResponse struct {
	AnonymizedText string           `json:"anonymized_text"`
	Entities       []DetectedEntity `json:"entities"`
	Error          string           `json:"error,omitempty"`
}

// DetectedEntity represents a detected entity in the text
type DetectedEntity struct {
	Type        string `json:"type"`
	Text        string `json:"text"`
	StartOffset int    `json:"start_offset"`
	EndOffset   int    `json:"end_offset"`
}

// NewAnonymizationService creates a new anonymization service
func NewAnonymizationService(baseURL string) *AnonymizationService {
	return &AnonymizationService{
		baseURL: baseURL,
	}
}

// AnonymizeText anonymizes sensitive information in text
func (s *AnonymizationService) AnonymizeText(text string) (string, error) {
	// Create request body
	reqBody := AnonymizeRequest{
		Text: text,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/api/v1/anonymize", s.baseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse response
	var result AnonymizeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for errors in the response
	if result.Error != "" {
		return "", fmt.Errorf("anonymization service error: %s", result.Error)
	}

	return result.AnonymizedText, nil
}
