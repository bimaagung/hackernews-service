package delivery

import (
	"hackernews-service/domain"
	"net/http"
)

type NewsHandler struct {
	NewsUseCase domain.NewsUsecase
}

func (h *NewsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	
	resTopStories, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	
	resItemStory, err := http.Get("https://hacker-news.firebaseio.com/v0/item/36073020.json?print=pretty")
	
	if err != nil {
		return 
	}

	defer resTopStories.Body.Close()

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.WriteHeader(http.StatusOK)
}