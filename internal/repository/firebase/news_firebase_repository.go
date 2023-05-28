package firebaserepository

import (
	"encoding/json"
	"fmt"
	"hackernews-service/constants"
	"hackernews-service/domain"
	mw "hackernews-service/internal/delivery/middleware"
	"io/ioutil"
)

func NewNewsFirebaseRepository(client domain.HTTPClient) domain.NewsFirebaseRepository {
	return &newsFirebaseRepository{
		HTTPClient: client,
	}
}

type newsFirebaseRepository struct {
	HTTPClient domain.HTTPClient
}

func (repository *newsFirebaseRepository) GetTopStories()([]int, error){
	var topStories []int

	resTopStories, err := repository.HTTPClient.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty") 
	
	if err != nil {
		return nil, &mw.NotFoundError{Message: constants.TopStoriesNotFound}
	}

	bodyTopStories, err := ioutil.ReadAll(resTopStories.Body)
	
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}

	err = json.Unmarshal(bodyTopStories, &topStories)
	
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}
	return topStories, nil
}

func (repository *newsFirebaseRepository) GetStoryById(id int) (*domain.Story, error){
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)

	res, err := repository.HTTPClient.Get(url)
	if err != nil {
		return nil, &mw.NotFoundError{Message: constants.StoryNotFound}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}

	var itemStory *domain.Story
	err = json.Unmarshal(body, &itemStory)
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}

	return itemStory, nil
}

func (repository *newsFirebaseRepository) GetCommentById(id int) (*domain.Comment, error){
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)

	res, err := repository.HTTPClient.Get(url)
	if err != nil {
		return nil, &mw.NotFoundError{Message: constants.CommentNotFound}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}

	var itemComment *domain.Comment
	err = json.Unmarshal(body, &itemComment)
	if err != nil {
		return nil, &mw.InternalServerError{Message: err.Error()}
	}

	return itemComment, nil
}