package di

import (
	"github.com/peya/src/domain/repos"
	"github.com/peya/src/infra/postgres"
	"github.com/peya/src/usecases"
)

func ProvideHealthCheckUsecaseImpl(r repos.HealthCheckRepo) *usecases.HealthCheckUsecaseImpl {
	return &usecases.HealthCheckUsecaseImpl{HealthCheckRepo: r}
}

func ProvideHealthCheckDBRepo() *infra.HealthCheckDBRepo {
	return &infra.HealthCheckDBRepo{}
}

func ProvideCreateReviewUsecaseImpl(r repos.ReviewRepo) *usecases.CreateReviewUsecaseImpl {
	return &usecases.CreateReviewUsecaseImpl{ReviewRepo: r}
}

func ProvideCreateReviewDBRepo() *infra.ReviewDBRepo {
	return &infra.ReviewDBRepo{}
}
