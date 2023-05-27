package domain

import "net/http"

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type NewsUsecase interface {
	GetAll(story *Item) ([]*Item, error)
}

type NewsFirebaseRepository interface {
	GetTopStories() ([]int, error)
	GetItemById() (*Item, error)
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}