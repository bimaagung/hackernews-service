package delivery

import (
	"encoding/json"
	"fmt"
	"hackernews-service/domain"
	"io/ioutil"
	"net/http"
)

type NewsHandler struct {
	NewsUseCase domain.NewsUsecase
}

func (h *NewsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var itemStories []*domain.Story
	var topStories []int

	resTopStories, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bodyTopStories, err := ioutil.ReadAll(resTopStories.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bodyTopStories, &topStories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a channel to receive itemStories
	itemStoryChan := make(chan *domain.Story)

	for _, v := range topStories {
		urlItemStory := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", v)

		go func(url string) {
			resItemStory, err := http.Get(url)
			if err != nil {
				itemStoryChan <- nil // Signal an error
				return
			}
			defer resItemStory.Body.Close()

			bodyItemStory, err := ioutil.ReadAll(resItemStory.Body)
			if err != nil {
				itemStoryChan <- nil // Signal an error
				return
			}

			var itemStory *domain.Story
			err = json.Unmarshal(bodyItemStory, &itemStory)
			if err != nil {
				itemStoryChan <- nil // Signal an error
				return
			}

			itemStoryChan <- itemStory // Send the itemStory through the channel
		}(urlItemStory)
	}

	// Wait for all goroutines to finish and collect the itemStories
	for range topStories {
		itemStory := <-itemStoryChan
		if itemStory != nil {
			itemStories = append(itemStories, itemStory)
		}
	}

	close(itemStoryChan)

	// Convert the stories to JSON
	response, err := json.Marshal(itemStories)
	if err != nil {
		// Handle the error and return an appropriate response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
