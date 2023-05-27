package domain

import "net/http"

type Story struct {
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

type ResStories struct {
	By          		string 	`json:"by"`
	Descendants 		int    	`json:"descendants"`
	ID          		int    	`json:"id"`
	TotalComment        int  	`json:"total_comment"`
	Score       		int    	`json:"score"`
	Time        		int    	`json:"time"`
	Title       		string 	`json:"title"`
	URL         		string 	`json:"url"`
}

type NewsUsecase interface {
	GetAll() ([]*ResStories, error)
}

type NewsFirebaseRepository interface {
	GetTopStories() ([]int, error)
	GetStoryById(id int) (*Story, error)
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}