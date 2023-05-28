package delivery

import (
	"encoding/json"
	"hackernews-service/domain"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type NewsHandler struct {
	NewsUseCase domain.NewsUsecase
}

func (h *NewsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	
	stories, err := h.NewsUseCase.GetAll()

	response, err := json.Marshal(stories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *NewsHandler) GetStoryById(w http.ResponseWriter, r *http.Request) {
	storyId := chi.URLParam(r, "storyId") 
	
	storyIdInt, err := strconv.Atoi(storyId)

	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	story, err := h.NewsUseCase.GetStoryById(storyIdInt)

	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	response, err := json.Marshal(story)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *NewsHandler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId") 
	
	commentIdInt, err := strconv.Atoi(commentId)

	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	story, err := h.NewsUseCase.GetCommentById(commentIdInt)

	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	response, err := json.Marshal(story)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
