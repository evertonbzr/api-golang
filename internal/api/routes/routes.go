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
	todoHandler := handler.NewTodoHandler(r.DB)

	app.Use(middlewares.DecodeJwtMw())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", authHandler.Login())
	app.Post("/register", authHandler.Register())
	app.Get("/me", userHandler.GetMe())

	app.Get("/users", userHandler.List())

	app.Route("/todo", func(api fiber.Router) {
		api.Post("/", todoHandler.Create()).Name("create")
		api.Get("/", todoHandler.List()).Name("list")
		api.Put("/:id", todoHandler.Update()).Name("update.put")
		api.Patch("/:id", todoHandler.Update()).Name("update.patch")
	}, "todo.")
}
