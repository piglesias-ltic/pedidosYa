package usecases

import (
	"github.com/peya/src/domain/entities"
	"github.com/peya/src/domain/repos"
)

type CreateReviewUsecase interface {
	Create(entities.Review) (*entities.Review, error)
}

type CreateReviewUsecaseImpl struct {
	ReviewRepo repos.ReviewRepo
}

func (c *CreateReviewUsecaseImpl) Create(entities.Review) (*entities.Review, error) {
	return nil, nil
}
