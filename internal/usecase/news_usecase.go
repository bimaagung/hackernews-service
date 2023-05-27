package usecase

import (
	"hackernews-service/domain"
)



func NewNewsUsecase(newsFirebaseRepository domain.NewsFirebaseRepository) domain.NewsUsecase {
	return &newsUsecase{
		NewsFirebaseRepository: newsFirebaseRepository,
	}
}

type newsUsecase struct{
	NewsFirebaseRepository domain.NewsFirebaseRepository
}

func (uc newsUsecase) GetAll() ([]*domain.ResStories, error) {
	var itemStories []*domain.ResStories

	topStories, err := uc.NewsFirebaseRepository.GetTopStories() 

	if err != nil {
		return nil, err
	}

	itemStoryChan := make(chan *domain.ResStories)

	for _, v := range topStories {
		

		go func(id int) {
			story, err := uc.NewsFirebaseRepository.GetStoryById(id)
			itemStory := &domain.ResStories{
				ID: story.ID,
				By: story.By,
				Descendants: story.Descendants,
				TotalComment: len(story.Kids),
				Score: story.Score,
				Time: story.Time,
				Title: story.Title,
				URL: story.URL,
			}
			
			if err != nil {
				itemStoryChan <- nil 
				return
			}

			itemStoryChan <- itemStory
		}(v)
	}

	for range topStories {
		itemStory := <-itemStoryChan
		if itemStory != nil {
			itemStories = append(itemStories, itemStory)
		}
	}

	close(itemStoryChan)

	return itemStories, nil
}
