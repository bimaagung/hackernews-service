package delivery

import (
	"encoding/json"
	"hackernews-service/domain"
	"net/http"
)

type NewsHandler struct {
	NewsUseCase domain.NewsUsecase
}

func (h *NewsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	
	stories, err := h.NewsUseCase.GetAll()

	// Convert the stories to JSON
	response, err := json.Marshal(stories)
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
