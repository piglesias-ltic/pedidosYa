package usecases

import (
	"github.com/peya/src/domain/entities"
	"github.com/peya/src/domain/repos"
)

type HealthCheckUsecase interface {
	GetStatus() (*entities.VersionInfo, error)
}

type HealthCheckUsecaseImpl struct {
	HealthCheckRepo repos.HealthCheckRepo
}

func (u *HealthCheckUsecaseImpl) GetStatus() (*entities.VersionInfo, error) {
	return u.HealthCheckRepo.GetVersion()
}
