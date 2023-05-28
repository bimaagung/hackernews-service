package firebaserepository

import (
	"encoding/json"
	"fmt"
	"hackernews-service/domain"
	"io/ioutil"
	"time"
)

const dbTimeout = time.Second * 3

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
		return nil, err
	}

	bodyTopStories, err := ioutil.ReadAll(resTopStories.Body)
	
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyTopStories, &topStories)
	
	if err != nil {
		return nil, err
	}
	return topStories, nil
}

func (repository *newsFirebaseRepository) GetStoryById(id int) (*domain.Story, error){
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)

	res, err := repository.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var itemStory *domain.Story
	err = json.Unmarshal(body, &itemStory)
	if err != nil {
		return nil, err
	}

	return itemStory, nil
}