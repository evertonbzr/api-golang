package api

import (
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type APIConfig struct {
	Cache *redis.Client
	Port  string
}

func Run(cfg *APIConfig) {
	slog.Info(fmt.Sprintf("API running on port %s", cfg.Port))
}
