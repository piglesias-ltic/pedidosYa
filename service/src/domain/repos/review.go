package repos

import "github.com/peya/src/domain/entities"

type ReviewRepo interface {
	Create(entities.Review) (*entities.Review, error)
}
