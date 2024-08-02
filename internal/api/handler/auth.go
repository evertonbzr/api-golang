package handler

import (
	"time"

	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service *service.UserService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		Service: service.NewUserService(db),
	}
}

func (h *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.LoginRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		user, err := h.Service.GetUserByEmail(data.Email)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		if user.Password != data.Password {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "invalid credentials"})
		}

		token, _ := service.NewAccessToken(service.UserClaims{
			Id: user.ID,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		})

		return c.Status(fiber.StatusOK).JSON(map[string]string{"token": token})
	}
}

func (h *AuthHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.RegisterRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		user, _ := h.Service.GetUserByEmail(data.Email)

		if user.ID != 0 {
			return c.Status(fiber.StatusConflict).JSON(map[string]string{"error": "user already exists"})
		}

		user = &model.User{
			FullName: data.FullName,
			Email:    data.Email,
			Password: data.Password,
		}

		if err := h.Service.CreateUser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(nil)
		}

		return c.Status(fiber.StatusCreated).JSON(nil)
	}
}
