package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthcheckResponse struct {
	Success bool `json:"success"`
}

func RpcHealthcheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("RPC healthcheck called")
	response := &HealthcheckResponse{Success: true}

	out, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error marshalling response: %v", err)
	}

	return string(out), nil
}
