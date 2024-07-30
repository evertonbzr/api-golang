package routes

import (
	"net/http"

	"github.com/evertonbzr/api-golang/internal/api/handler"
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
	authHandler := handler.NewAuthHandler(r.DB)
	userHandler := handler.NewUserHandler(r.DB)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Post("/login", authHandler.Login())
	router.Post("/register", authHandler.Register())

	router.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.GetUserById())
	})
}
