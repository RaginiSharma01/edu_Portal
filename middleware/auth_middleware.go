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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		//  Check proper format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization format",
			})
		}

		// extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// verify JWT
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		//  safe type assertions
		userID, ok := claims["empId"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user id in token",
			})
		}

		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)

		// store in context
		c.Locals("userID", userID)
		c.Locals("email", email)
		c.Locals("role", role)

		return c.Next()
	}
}
