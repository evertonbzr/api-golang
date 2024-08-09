package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/evertonbzr/api-golang/internal/api"
	"github.com/evertonbzr/api-golang/internal/config"
	"github.com/evertonbzr/api-golang/pkg/infra"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	ctx := context.Background()

	infra.SetupDependecies(ctx, config.REDIS_URL, config.DATABASE_URL)
	// slog.Info("Connected to database")

	// queueConfig := &queue.QueueConfig{
	// 	URI:  config.NATS_URI,
	// 	Name: config.NAME,
	// }

	// nc, js, _ := queue.Start(queueConfig)

	apiCfg := &api.APIConfig{
		Port: config.PORT,
	}

	api.Start(apiCfg)
}
