package handler

import (
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Service: service.NewUserService(db),
	}
}

func (h *UserHandler) GetMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(int)

		return c.JSON(fiber.Map{
			"userId": userId,
		})
	}
}

func (h *UserHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := h.Service.ListUsers()

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"users": users,
		})
	}
}
