package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(uri string) *redis.Client {
	url, err := redis.ParseURL(uri)

	if err != nil {
		log.Fatal("Error parsing redis url", "error", err)
	}

	client := redis.NewClient(url)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Error connecting to redis", "error", err)
	}

	return client
}
