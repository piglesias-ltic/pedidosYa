package repos

import "github.com/peya/src/domain/entities"

type HealthCheckRepo interface {
	GetVersion() (*entities.VersionInfo, error)
}
