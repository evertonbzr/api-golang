package handler

import (
	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/evertonbzr/api-golang/internal/util"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserRepo: repository.NewUserRepository(),
	}
}

func (h *UserHandler) GetMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("userClaims").(*util.ModuleClaims)

		if !ok {
			return c.Status(500).JSON(fiber.Map{
				"error": "UserClaimsNotFound",
			})
		}

		return c.JSON(fiber.Map{
			"user": claims.User,
		})
	}
}

func (h *UserHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := h.UserRepo.ListNotAdminUsers()

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
