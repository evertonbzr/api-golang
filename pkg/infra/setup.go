package infra

import (
	"context"
	"log/slog"

	"github.com/evertonbzr/api-golang/pkg/infra/db"
	"github.com/evertonbzr/api-golang/pkg/infra/redis"
)

func SetupDependecies(ctx context.Context, redisURI string, postgresURI string) {
	db.New(postgresURI)
	slog.Info("Connected to database")

	redis.InitRedisClient(redisURI)
	slog.Info("Connected to redis")
}

func CleanUpDependecies() {
	db.Disconnect()
	slog.Info("Disconnected from database")

	redis.Disconnect()
	slog.Info("Disconnected from redis")
}
