package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedisClient(uri string) *redis.Client {
	url, err := redis.ParseURL(uri)
	if err != nil {
		log.Fatal("redis", err.Error())
	}

	if client != nil {
		slog.Warn("Redis client already initialized, returning the same instance")
		return client
	}

	client = redis.NewClient(url)

	return client
}

func GetRedisClient() *redis.Client {
	if client == nil {
		panic("Redis client not initialized")
	}
	return client
}

func Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = GetRedisClient().Set(ctx, key, b, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func HasKey(ctx context.Context, key string) bool {
	b, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if len(b) == 0 {
		return false
	}

	return true
}

func Get(ctx context.Context, key string, value interface{}) error {
	b, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return errors.New("KeyNotFound")
	}

	err = json.Unmarshal([]byte(b), value)

	if err != nil {
		return err
	}

	return nil
}

func Disconnect() {
	if client == nil {
		slog.Warn("redis No redis client found to disconect")
		return
	}

	err := client.Close()
	if err != nil {
		log.Fatal("redis", err.Error())
	}
}
