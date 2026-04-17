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
	dashboardHandler *handler.DashboardHandler,
	TeacherDashboard *handler.TeacherDashboardHandler,
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
	onboarding.Post("/forgot-password", userHandler.ForgotPassword)
	onboarding.Post("/reset-password", userHandler.ResetPassword)

	// classroom routes
	// admin only — create, manage, delete classrooms
	classroom := v1.Group("/classrooms", middleware.AdminOnly())
	classroom.Post("/create", classroomHandler.CreateClassroom)
	classroom.Get("/get", classroomHandler.GetClassrooms)
	classroom.Delete("/:id", classroomHandler.DeleteClassroom)
	classroom.Post("/:id/students", classroomHandler.AddStudentsToClassroom)
	classroom.Delete("/:id/students/:studentId", classroomHandler.RemoveStudentFromClassroom)

	// admin + teacher — view classroom detail
	classroomView := v1.Group("/classrooms")
	classroomView.Get("/:id", middleware.AdminOnly(), classroomHandler.GetClassroomByID)
	classroomView.Get("/:id", middleware.TeacherOnly(), classroomHandler.GetClassroomByID) // duplicate route — see note below
	// event routes
	event := v1.Group("/events", middleware.AdminOnly())
	event.Post("/create", eventHandler.CreateEvent)
	event.Get("/get", eventHandler.GetEvents)

	// salary routes
	salary := v1.Group("/salaries", middleware.AdminOnly())
	salary.Post("/create", salaryHandler.CreateSalary)
	salary.Get("/get", salaryHandler.GetAllSalaries)
	salary.Put("/:teacherId", salaryHandler.UpdateSalary)

	// timetable routes
	timetable := v1.Group("/timetable", middleware.AdminOnly(), middleware.TeacherOnly())
	timetable.Post("/create", timetableHandler.CreateTimetable)
	timetable.Get("/get", timetableHandler.GetTimetable)

	// marksheet routes
	marksheet := v1.Group("/marksheet")
	marksheet.Post("/create", marksheetHandler.CreateMarks)
	marksheet.Get("/get", marksheetHandler.GetMarks)
	marksheet.Get("/marks/pdf", marksheetHandler.DownloadMarksPDF)
	marksheet.Get("/:id/pdf", marksheetHandler.DownloadStudentPDF)

	// fetch routes
	getUsers := v1.Group("/fetch")
	getUsers.Get("/teachers", userHandler.GetAllTeachers)
	getUsers.Get("/students", userHandler.GetAllStudents)

	// student management — admin only
	students := v1.Group("/students", middleware.AdminOnly())
	students.Delete("/:id", userHandler.DeleteStudent)
	students.Put("/:id", userHandler.UpdateStudent)
	students.Patch("/:id/block", userHandler.BlockStudent)

	// teacher management — admin only
	teachers := v1.Group("/teachers", middleware.AdminOnly())
	teachers.Delete("/:id", userHandler.DeleteTeacher)
	teachers.Put("/:id", userHandler.UpdateTeacher)

	dashboard := v1.Group("/dashboard", middleware.AdminOnly())
	dashboard.Get("/admin", dashboardHandler.GetAdminDashboard)

	//teacher dashboard

	dashboardTeacher := v1.Group("/dashboard")
	dashboardTeacher.Post("/teacher", TeacherDashboard.GetTeacherDashboard)
}
