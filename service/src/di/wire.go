//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/peya/src/domain/repos"
	"github.com/peya/src/infra/postgres"
	"github.com/peya/src/usecases"
)

var Set = wire.NewSet(
	ProvideHealthCheckUsecaseImpl,
	ProvideHealthCheckDBRepo,
	wire.Bind(new(repos.HealthCheckRepo), new(*infra.HealthCheckDBRepo)),
	wire.Bind(new(usecases.HealthCheckUsecase), new(*usecases.HealthCheckUsecaseImpl)),
	wire.Struct(new(Injector), "*"),
)

type Injector struct {
	HealthCheckUsecase usecases.HealthCheckUsecase
}

func InitializeInjector() *Injector {
	wire.Build(Set)
	return &Injector{}
}
