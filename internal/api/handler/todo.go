package handler

import (
	"strconv"

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
		data := types.CreateOrUpdateTodoRequest{}

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

func (h *TodoHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		todos, err := h.Service.List()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"todos": todos,
		})
	}
}

func (h *TodoHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoId := c.Params("id")

		var data types.CreateOrUpdateTodoRequest
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		id64, err := strconv.Atoi(todoId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}
		id := uint(id64)

		todo, err := h.Service.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}

		if data.Title != "" {
			todo.Title = data.Title
		}
		if data.Description != "" {
			todo.Description = data.Description
		}
		if data.Status != "" {
			todo.Status = data.Status
		}

		if err := h.Service.Update(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Todo updated successfully",
			"todo":    todo,
		})
	}
}
