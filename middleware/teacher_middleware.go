package middleware

import (
	"smp/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func TeacherOnly() fiber.Handler {
	return func(c fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "token required",
			})
		}

		parts := strings.Split(authHeader, "")
		if len(parts) != 2 {
			return c.Status(401).JSON(fiber.Map{
				"error": "inavlid format add the bearer",
			})
		}

		token := parts[1]
		claims, err := utils.VerifyJWT(token)
		if err != nil {
			return err
		}
		role := claims["role"]
		if role == nil || role.(string) != "teacher" {
			return c.Status(403).JSON(fiber.Map{
				"error": "only admin allowed",
			})
		}

		return c.Next()
	}
}
