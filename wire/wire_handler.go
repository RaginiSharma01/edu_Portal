package wire

import "smp/handler"

type Handlers struct {
	UserHandler             *handler.UserHandler
	ClassroomHandler        *handler.ClassroomHandler
	EventHandler            *handler.EventHandler
	SalaryHandler           *handler.SalaryHandler
	TimetableHandler        *handler.TimetableHandler
	MarksHandler            *handler.MarksHandler
	DashboardHandler        *handler.DashboardHandler
	TeacherDashboardHandler *handler.TeacherDashboardHandler
}
