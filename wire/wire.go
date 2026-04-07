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

func InitializeUserHandler() *handler.UserHandler {

	wire.Build(
		config.LoadConfig,
		db.ProvidePgDb,
		db.ProvidePool,
		db.ProvideRedis,

		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
	)

	return &handler.UserHandler{}
}
