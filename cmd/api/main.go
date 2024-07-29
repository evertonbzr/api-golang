package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/api-golang/internal/api"
	"github.com/evertonbzr/api-golang/internal/cache"
	"github.com/evertonbzr/api-golang/internal/config"
	"github.com/evertonbzr/api-golang/internal/queue"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	cache := cache.InitRedis(config.REDIS_URL)
	slog.Info("Connected to redis")

	queueConfig := &queue.QueueConfig{
		URI:  config.NATS_URI,
		Name: config.NAME,
	}

	nc, js, _ := queue.Start(queueConfig)

	apiCfg := &api.APIConfig{
		Cache:          cache,
		Port:           config.PORT,
		NatsConnection: nc,
		JetStream:      js,
	}

	api.Start(apiCfg)
}
