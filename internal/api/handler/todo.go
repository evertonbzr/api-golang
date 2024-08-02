package handler

import (
	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{
		Service: service.NewTodoService(db),
	}
}

func (h *TodoHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data := types.CreateTodoRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		todo := model.Todo{
			Title:       data.Title,
			Description: data.Description,
			Status:      data.Status,
			UserID:      1,
		}

		if err := h.Service.Create(
			[]model.Todo{todo},
		); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"message": "Create",
		})
	}
}
