package routes

import (
	"github.com/evertonbzr/api-golang/internal/api/handler"
	"github.com/evertonbzr/api-golang/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
}

func NewRoute() *RouteConfig {
	return &RouteConfig{}
}

func (r *RouteConfig) SetRoutesFiber(app *fiber.App) {
	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()
	bookHandler := handler.NewBookHandler()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", authHandler.Login())
	app.Post("/register", authHandler.Register())

	// app.Use(middlewares.AuthJwtMw("ad"))
	app.Get("/me", middlewares.AuthJwtMw("admin"), userHandler.GetMe())
	app.Get("/users", middlewares.AuthJwtMw("admin"), userHandler.List())

	app.Route("/books", func(api fiber.Router) {
		api.Post("/", bookHandler.Create()).Name("create")
		api.Get("/", bookHandler.List()).Name("list")
		api.Put("/:id", bookHandler.Update()).Name("update.put")
		api.Patch("/:id", bookHandler.Update()).Name("update.patch")
	}, "todo.")
}
