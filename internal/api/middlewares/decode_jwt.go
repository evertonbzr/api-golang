package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DecodeJwtMw() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()

		c.Locals("userId", 2)

		jwtTokenBearer := headers["X-Jwt-Token"]

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

		return c.Next()
	}
}
