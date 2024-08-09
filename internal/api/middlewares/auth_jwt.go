package middlewares

import (
	"log/slog"
	"strings"

	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/evertonbzr/api-golang/internal/util"
	"github.com/evertonbzr/api-golang/pkg/infra/redis"
	"github.com/gofiber/fiber/v2"
)

func createUserContextAndCache(ctx *fiber.Ctx, jwt string) (*util.ModuleClaims, error) {
	claims := &util.ModuleClaims{}
	err := redis.Get(ctx.Context(), jwt, claims)

	if err == nil {
		slog.Info("A redis cache was found", "claims", claims)
		ctx.Locals("userClaims", claims)
		return nil, nil
	}

	token, claims, err := util.DecodeJWT(jwt)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository()

	updatedUser, err := userRepo.GetUserById(claims.User.ID)
	if err != nil {
		return nil, err
	}

	cacheDuration, err := util.GetDurationFromJWT(token)
	if err != nil {
		return nil, err
	}

	claims.User = *updatedUser

	err = redis.Save(ctx.Context(), jwt, claims, cacheDuration)
	if err != nil {
		return nil, err
	}

	slog.Info("[Authorization]", "jwt=", token.Raw, "user=", claims.User, "expDuration=", cacheDuration)
	ctx.Locals("userClaims", claims)
	return claims, nil
}

func AuthJwtMw(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		var jwtTokenBearer []string = make([]string, 0)

		jwtTokenBearer = headers["X-Jwt-Token"]
		if len(jwtTokenBearer) == 0 {
			jwtTokenBearer = headers["Authorization"]
		}
		if len(jwtTokenBearer) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		jwtToken := jwtTokenBearer[0]
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)

		if jwtToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		_, err := createUserContextAndCache(c, jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "error": err.Error()})
		}

		if len(roles) > 0 {
			claims := c.Locals("userClaims").(*util.ModuleClaims)
			for _, role := range roles {
				if role == claims.User.Role {
					return c.Next()
				}
			}

			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden", "error": "You don't have permission to access this resource"})

		}

		return c.Next()
	}
}
