package middleware

import (
	"strings"

	"smp/utils"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		// Expect: Bearer <token>
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		c.Locals("userID", claims["empId"])
		c.Locals("email", claims["email"])

		return c.Next()
	}
}
