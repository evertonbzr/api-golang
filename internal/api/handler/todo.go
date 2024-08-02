package handler

import (
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (h *TodoHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Create",
		})
	}
}
