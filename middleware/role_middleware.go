package middleware

import (
	"smp/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AdminOnly() fiber.Handler {
	return func(c fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing token",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		token := parts[1]

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		role := claims["role"]

		if role == nil || role.(string) != "admin" {
			return c.Status(403).JSON(fiber.Map{
				"error": "only admin allowed",
			})
		}

		return c.Next()
	}
}
