package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/api-golang/internal/api"
	"github.com/evertonbzr/api-golang/internal/cache"
	"github.com/evertonbzr/api-golang/internal/config"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	cache := cache.InitRedis(config.REDIS_URL)
	slog.Info("Connected to redis")

	cfg := &api.APIConfig{
		Cache: cache,
		Port:  config.PORT,
	}

	api.Run(cfg)
}
