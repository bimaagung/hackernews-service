package usecase

import (
	"hackernews-service/domain"
	memorycache "hackernews-service/pkg/memory_cache"
	"strconv"
	"time"
)

const (
	cacheExpirationTime = 5 * time.Minute
	cacheUpdateInterval = 30 * time.Second
)

func NewNewsUsecase(newsFirebaseRepository domain.NewsFirebaseRepository, cache memorycache.Cache) domain.NewsUsecase {
	return &newsUsecase{
		NewsFirebaseRepository	: newsFirebaseRepository,
		Cache					:cache,
	}
}

type newsUsecase struct{
	NewsFirebaseRepository domain.NewsFirebaseRepository
	Cache                  memorycache.Cache
}

func (uc *newsUsecase)StartCacheUpdate() {
	go uc.updateCachePeriodically()
}

func (uc *newsUsecase) updateCachePeriodically() {
	for {
		uc.GetAll()
		time.Sleep(cacheUpdateInterval)
	}
}

func (uc *newsUsecase) GetAll() ([]*domain.ResStories, error) {
	
	itemStoriesCache, err := uc.Cache.Get("all_stories")

	if err == nil {
		if stories, ok := itemStoriesCache.([]*domain.ResStories); ok {
			return stories, nil
		}
	}

	topStories, err := uc.NewsFirebaseRepository.GetTopStories()
	if err != nil {
		return nil, err
	}

	itemStoryChan := make(chan *domain.ResStories)

	for _, id := range topStories {
		storyRaw, err := uc.Cache.Get("story_" + strconv.Itoa(id))
		if err == nil {
			if story, ok := storyRaw.(*domain.Story); ok {
				itemStoryChan <- &domain.ResStories{
					ID:            story.ID,
					By:            story.By,
					Descendants:   story.Descendants,
					TotalComment:  len(story.Kids),
					Score:         story.Score,
					Time:          story.Time,
					Title:         story.Title,
					URL:           story.URL,
				}
				continue
			}
		}

		go func(storyID int) {
			story, err := uc.NewsFirebaseRepository.GetStoryById(storyID)
			itemStory := &domain.ResStories{
				ID:            story.ID,
				By:            story.By,
				Descendants:   story.Descendants,
				TotalComment:  len(story.Kids),
				Score:         story.Score,
				Time:          story.Time,
				Title:         story.Title,
				URL:           story.URL,
			}

			if err != nil {
				itemStoryChan <- nil
				return
			}

			uc.Cache.Set("story_"+strconv.Itoa(storyID), story, cacheExpirationTime)

			itemStoryChan <- itemStory
		}(id)
	}


	var itemStories []*domain.ResStories

	for range topStories {
		itemStory := <-itemStoryChan
		if itemStory != nil {
			itemStories = append(itemStories, itemStory)
		}
	}

	close(itemStoryChan)

	uc.Cache.Set("all_stories", itemStories, cacheExpirationTime)

	return itemStories, nil
}

func (uc *newsUsecase) GetStoryById(id int) (*domain.ResStory, error) {
	var itemComments []*domain.ResComment
	var resStory *domain.ResStory

	story, err := uc.NewsFirebaseRepository.GetStoryById(id)

	if err != nil {
		return nil, err
	}

	itemCommentChan := make(chan *domain.ResComment)

	for _, v := range story.Kids {
		

		go func(id int) {
			comment, err := uc.NewsFirebaseRepository.GetCommentById(id)
			resComment := &domain.ResComment{
				ID: comment.ID,
				By: comment.By,
				Parent: comment.Parent,
				Text: comment.Text,
				Time: comment.Time,
				Type: comment.Type,
			}
			
			if err != nil {
				itemCommentChan <- nil 
				return
			}

			itemCommentChan <- resComment
		}(v)
	}

	for range story.Kids {
		itemComment := <-itemCommentChan
		if itemComment != nil {
			itemComments = append(itemComments, itemComment)
		}
	}

	close(itemCommentChan)

	resStory = &domain.ResStory{
		ID: story.ID,
		By: story.By,
		Descendants: story.Descendants,
		Comments: itemComments,
		Score: story.Score,
		Title: story.Title,
		Time: story.Time,
		URL: story.URL,
		Type: story.Type,
	}

	return resStory, nil
}

func (uc *newsUsecase) GetCommentById(id int)(*domain.ResComment, error) {
	var itemComments []*domain.ResComment
	var resComment *domain.ResComment

	comment, err := uc.NewsFirebaseRepository.GetCommentById(id)

	if err != nil {
		return nil, err
	}

	itemCommentChan := make(chan *domain.ResComment)

	for _, v := range comment.Kids {
		

		go func(id int) {
			comment, err := uc.NewsFirebaseRepository.GetCommentById(id)
			resComment := &domain.ResComment{
				ID: comment.ID,
				By: comment.By,
				Parent: comment.Parent,
				Text: comment.Text,
				Time: comment.Time,
				Type: comment.Type,
			}
			
			if err != nil {
				itemCommentChan <- nil 
				return
			}

			itemCommentChan <- resComment
		}(v)
	}

	for range comment.Kids {
		itemComment := <-itemCommentChan
		if itemComment != nil {
			itemComments = append(itemComments, itemComment)
		}
	}

	close(itemCommentChan)

	resComment = &domain.ResComment{
		ID: comment.ID,
		By: comment.By,
		Comments: itemComments,
		Parent: comment.Parent,
		Text: comment.Text,
		Time: comment.Time,
		Type: comment.Type,
	}

	return resComment, nil
}