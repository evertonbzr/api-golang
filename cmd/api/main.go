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

	infra.SetupDependecies(ctx, &infra.InfraConfig{
		RedisURI:    config.REDIS_URL,
		PostgresURI: config.DATABASE_URL,
		NatsURI:     config.NATS_URI,
		NatsName:    config.NAME,
	})

	apiCfg := &api.APIConfig{
		Port: config.PORT,
	}

	api.Start(apiCfg)
}
