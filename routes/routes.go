package routes

import (
	"smp/handler"
	"smp/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
	classroomHandler *handler.ClassroomHandler,
	eventHandler *handler.EventHandler,
	salaryHandler *handler.SalaryHandler,
	timetableHandler *handler.TimetableHandler,
	marksheetHandler *handler.MarksHandler,

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
	classroom := v1.Group("/classrooms", middleware.AdminOnly())

	classroom.Post("/create", classroomHandler.CreateClassroom)
	classroom.Get("/get", classroomHandler.GetClassrooms)

	// event routes
	event := v1.Group("/events", middleware.AdminOnly())

	event.Post("/create", eventHandler.CreateEvent)
	event.Get("/get", eventHandler.GetEvents)

	// salary routes
	salary := v1.Group("/salaries", middleware.AdminOnly())

	salary.Post("/create", salaryHandler.CreateSalary)
	salary.Get("/get", salaryHandler.GetAllSalaries)

	//time table

	timetable := v1.Group("/timetable", middleware.AdminOnly())

	timetable.Post("/create", timetableHandler.CreateTimetable)
	timetable.Get("/get", timetableHandler.GetTimetable)


	// marksheet

	marksheet := v1.Group("/marksheet")
	marksheet.Post("/create" , marksheetHandler.CreateMarks)
	marksheet.Get("/get" , marksheetHandler.GetMarks)
}
