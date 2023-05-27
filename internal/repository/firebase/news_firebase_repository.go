package firebaserepository

import (
	"encoding/json"
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

func (repository *newsFirebaseRepository) GetItemById() (*domain.Item, error){
	return nil, nil
}