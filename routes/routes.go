package routes

import (
	"smp/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler, //classroomHandler *handler.ClassroomHandler
) {

	api := app.Group("/api")
	smp := api.Group("/smp")
	v1 := smp.Group("/v1")

	// onboarding routes
	onboarding := v1.Group("/onboarding")
	signup := onboarding.Group("/signup")

	signup.Post("/teacher", userHandler.OnboardTeacher)
	signup.Post("/student", userHandler.OnboardStudent)
	onboarding.Post("/verify-otp", userHandler.VerifyOTP)
	onboarding.Post("/login", userHandler.Login)

	// classroom routes
	//classroom := v1.Group("/classrooms")

	//classroom.Post("/", classroomHandler.CreateClassroom)

}
