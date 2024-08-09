package handler

import (
	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/evertonbzr/api-golang/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		UserRepo: repository.NewUserRepository(),
	}
}

func (h *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.LoginRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		user, err := h.UserRepo.GetUserByEmail(data.Email)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email or password incorrect"})
		}

		if user.Password != data.Password {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "email or password incorrect"})
		}

		token, _ := util.GenerateJwt(user)

		return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"token": token, "user": user})
	}
}

func (h *AuthHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.RegisterRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		user, _ := h.UserRepo.GetUserByEmail(data.Email)

		if user.ID != 0 {
			return c.Status(fiber.StatusConflict).JSON(map[string]string{"error": "user already exists"})
		}

		user = &model.User{
			FullName: data.FullName,
			Email:    data.Email,
			Password: data.Password,
			Role:     "user",
		}

		if err := h.UserRepo.CreateUser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(nil)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
	}
}
