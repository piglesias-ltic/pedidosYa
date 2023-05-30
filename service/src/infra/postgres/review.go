package infra

import "github.com/peya/src/domain/entities"

type ReviewDBRepo struct{}

func (r ReviewDBRepo) Create(entities.Review) (*entities.Review, error) {
	return nil, nil
}
