package routes

import (
	"smp/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	api := app.Group("/api")
	smp := api.Group("/smp")
	v1 := smp.Group("/v1")

	onboarding := v1.Group("/onboarding")
	signup := onboarding.Group("/signup")

	signup.Post("/teacher", userHandler.OnboardTeacher)

}
