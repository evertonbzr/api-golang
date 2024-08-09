package routes

import (
	"github.com/evertonbzr/api-golang/internal/api/handler"
	"github.com/evertonbzr/api-golang/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
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

func (r *RouteConfig) SetRoutesFiber(app *fiber.App) {
	authHandler := handler.NewAuthHandler(r.DB)
	userHandler := handler.NewUserHandler(r.DB)
	bookHandler := handler.NewBookHandler(r.DB)

	app.Use(middlewares.DecodeJwtMw())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", authHandler.Login())
	app.Post("/register", authHandler.Register())
	app.Get("/me", userHandler.GetMe())

	app.Get("/users", userHandler.List())

	app.Route("/books", func(api fiber.Router) {
		api.Post("/", bookHandler.Create()).Name("create")
		api.Get("/", bookHandler.List()).Name("list")
		api.Put("/:id", bookHandler.Update()).Name("update.put")
		api.Patch("/:id", bookHandler.Update()).Name("update.patch")
	}, "todo.")
}
