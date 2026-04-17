//go:build wireinject
// +build wireinject

package wire

import (
	"smp/config"
	"smp/db"
	"smp/handler"
	"smp/repository"
	"smp/service"

	"github.com/google/wire"
)

func InitializeHandlers() *Handlers {
	wire.Build(
		config.LoadConfig,
		db.ProvidePgDb,
		db.ProvidePool,
		db.ProvideRedis,

		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,

		repository.NewClassroomRepo,
		service.NewClassroomService,
		handler.NewClassroomHandler,

		repository.NewEventRepository,
		service.NewEventService,
		handler.NewEventHandler,

		repository.NewSalaryRepository,
		service.NewSalaryService,
		handler.NewSalaryHandler,

		repository.NewTimetableRepository,
		service.NewTimetableService,
		handler.NewTimetableHandler,

		repository.NewMarksRepository,
		service.NewMarksService,
		handler.NewMarksHandler,

		repository.NewDashboardRepository,
		service.NewDashboardService,
		handler.NewDashboardHandler,

		repository.NewTeacherDashboardRepository,
		service.NewTeacherDashboardService,
		handler.NewTeacherDashboardHandler,

		wire.Struct(new(Handlers), "*"),
	)

	return &Handlers{}
}
