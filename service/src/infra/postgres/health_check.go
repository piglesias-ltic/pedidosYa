package infra

import (
	"github.com/peya/src/domain/entities"
)

type HealthCheckDBRepo struct{}

func (r *HealthCheckDBRepo) GetVersion() (*entities.VersionInfo, error) {
	db := getConnection()
	defer db.Close()

	versionInfo := entities.VersionInfo{}
	if err := db.QueryRow(
		"SELECT app_version, release_date FROM main.version",
	).Scan(&versionInfo.Version, &versionInfo.ReleaseDate); err != nil {
		return nil, err
	}

	return &versionInfo, nil
}
