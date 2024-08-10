package infra

import (
	"context"
	"log/slog"

	"github.com/evertonbzr/api-golang/pkg/infra/db"
	"github.com/evertonbzr/api-golang/pkg/infra/queue"
	"github.com/evertonbzr/api-golang/pkg/infra/redis"
)

type InfraConfig struct {
	RedisURI    string
	PostgresURI string
	NatsURI     string
	NatsName    string
}

func SetupDependecies(ctx context.Context, cfg *InfraConfig) {
	db.New(cfg.PostgresURI)
	slog.Info("Connected to database")

	redis.InitRedisClient(cfg.RedisURI)
	slog.Info("Connected to redis")

	queue.InitNatsClient(ctx, cfg.NatsURI, cfg.NatsName)
	slog.Info("Connected to nats")
}

func CleanUpDependecies() {
	db.Disconnect()
	slog.Info("Disconnected from database")

	redis.Disconnect()
	slog.Info("Disconnected from redis")

	queue.CloseNatsConnection()
	slog.Info("Disconnected from nats")
}
