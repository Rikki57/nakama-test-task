package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/heroiclabs/nakama-common/runtime"
)

type Payload struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

type Response struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

func RpcCountHash(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Got a call to RpcCountHash with payload: %v", payload)
	var unmarshalledPayload Payload
	if err := json.Unmarshal([]byte(payload), &unmarshalledPayload); err != nil {
		logger.Error("Error unmarshalling payload: %v", err)
		return "", err
	}

	setDefaultValues(&unmarshalledPayload)

	filePath := filepath.Join("/nakama/data/files/", unmarshalledPayload.Type, unmarshalledPayload.Version+".json")
	content, err := readFile(filePath, logger)
	if err != nil {
		return "", err
	}

	calculatedHash := calculateHash(content)
	logger.Debug("Calculated Hash: %s", calculatedHash)

	responseContent := determineResponseContent(unmarshalledPayload.Hash, calculatedHash, content, logger)

	if err := saveToDatabase(ctx, db, unmarshalledPayload, calculatedHash, responseContent, logger); err != nil {
		return "", err
	}

	response := Response{
		Type:    unmarshalledPayload.Type,
		Version: unmarshalledPayload.Version,
		Hash:    calculatedHash,
		Content: responseContent,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error marshalling response: %v", err)
		return "", err
	}

	return string(responseBytes), nil
}

func setDefaultValues(payload *Payload) {
	if payload.Type == "" {
		payload.Type = "core"
	}
	if payload.Version == "" {
		payload.Version = "1.0.0"
	}
}

func readFile(filePath string, logger runtime.Logger) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error("Error opening file: %s", filePath)
		return nil, errors.New("file not found")
	}
	return content, nil
}

func calculateHash(content []byte) string {
	hash := sha256.New()
	hash.Write(content)
	return hex.EncodeToString(hash.Sum(nil))
}

func determineResponseContent(payloadHash, calculatedHash string, content []byte, logger runtime.Logger) string {
	logger.Debug("unmarshalledPayload Hash: '%s'", payloadHash)
	if payloadHash == "" || payloadHash != calculatedHash {
		return ""
	}
	return string(content)
}

func saveToDatabase(ctx context.Context, db *sql.DB, payload Payload, calculatedHash string, responseContent string, logger runtime.Logger) error {
	query := `INSERT INTO files (type, version, hash, content) VALUES ($1, $2, $3, $4)`
	params := []interface{}{payload.Type, payload.Version, calculatedHash, responseContent}
	if _, err := db.ExecContext(ctx, query, params...); err != nil {
		logger.Error("Error executing DB request, query: %s", query)
		return err
	}

	return nil
}
