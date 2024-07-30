package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RouteConfig struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewRoute(db *gorm.DB, cache *redis.Client) *RouteConfig {
	return &RouteConfig{
		DB:    db,
		Cache: cache,
	}
}

func (r *RouteConfig) SetRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
}
