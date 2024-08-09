package middlewares

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/evertonbzr/api-golang/internal/util"
	"github.com/evertonbzr/api-golang/pkg/infra/redis"
	"github.com/gofiber/fiber/v2"
)

func createUserContextAndCache(ctx *fiber.Ctx, jwt string) error {
	claims := &util.ModuleClaims{}
	err := redis.Get(ctx.Context(), jwt, claims)

	if err == nil {
		slog.Info("A redis cache was found", "claims", claims)
		ctx.Locals("user", claims)
		return nil
	}

	token, claims, err := util.DecodeJWT(jwt)
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepository()

	updatedUser, err := userRepo.GetUserById(claims.User.ID)
	if err != nil {
		return err
	}

	cacheDuration, err := util.GetDurationFromJWT(token)
	if err != nil {
		return err
	}

	claims.User = *updatedUser

	err = redis.Save(ctx.Context(), jwt, claims, cacheDuration)
	if err != nil {
		return errors.New("ErrorSavingCache")
	}

	return nil
}

func AuthJwtMw() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		var jwtTokenBearer []string = make([]string, 0)

		jwtTokenBearer = headers["X-Jwt-Token"]
		if len(jwtTokenBearer) == 0 {
			jwtTokenBearer = headers["Authorization"]
		}
		if len(jwtTokenBearer) == 0 {
			return c.Next()
		}

		jwtToken := jwtTokenBearer[0]
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)

		if jwtToken == "" {
			return c.Next()
		}

		err := createUserContextAndCache(c, jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "error": err.Error()})
		}

		return c.Next()
	}
}
