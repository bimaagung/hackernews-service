package usecase

import "hackernews-service/domain"

type newsUsecase struct{}

func NewNewsUsecase() domain.NewsUsecase {
	return &newsUsecase{}
}

func (u newsUsecase) GetAll(story *domain.Story) (stories []*domain.Story, err error) {
	return nil, nil
}
