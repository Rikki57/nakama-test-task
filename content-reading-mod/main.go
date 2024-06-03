package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()
	err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)
	if err != nil {
		return err
	}
	err2 := initializer.RegisterRpc("countHash", RpcCountHash)
	if err2 != nil {
		return err2
	}
	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())
	return nil
}
