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

	app.Get("/me", middlewares.AuthJwtMw("admin"), userHandler.GetMe())
	app.Get("/users", middlewares.AuthJwtMw("admin"), userHandler.List())

	app.Route("/books", func(api fiber.Router) {
		api.Get("/", middlewares.AuthJwtMw(), bookHandler.List()).Name("list")
		api.Use(middlewares.AuthJwtMw("admin"))
		api.Post("/", bookHandler.Create()).Name("create")
		api.Put("/:id", bookHandler.Update()).Name("update.put")
		api.Patch("/:id", bookHandler.Update()).Name("update.patch")
	}, "todo.")

	app.Route("/borrowings", func(api fiber.Router) {
		api.Use(middlewares.AuthJwtMw("admin"))
		api.Get("/", handler.NewBorrowingHandler().List())
		api.Put("/", handler.NewBorrowingHandler().Set())
	})
}
