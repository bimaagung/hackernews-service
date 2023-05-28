package delivery

import (
	"hackernews-service/constants"
	"hackernews-service/domain"
	"hackernews-service/helpers/response"
	mw "hackernews-service/internal/delivery/middleware"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type NewsHandler struct {
	NewsUseCase domain.NewsUsecase
}

func (h *NewsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	
	stories, err := h.NewsUseCase.GetAll()

	if err != nil {
		panic(err)
	}

	response.Success(w, http.StatusOK, stories)
}

func (h *NewsHandler) GetStoryById(w http.ResponseWriter, r *http.Request) {
	storyId := chi.URLParam(r, "storyId") 
	
	storyIdInt, err := strconv.Atoi(storyId)

	if err != nil {
		panic(&mw.BadRequestError{Message: constants.IdNotInt})
	}

	story, err := h.NewsUseCase.GetStoryById(storyIdInt)

	if err != nil {
		panic(err)
	}

	response.Success(w, http.StatusOK, story)
}

func (h *NewsHandler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId") 
	
	commentIdInt, err := strconv.Atoi(commentId)

	if err != nil {
		panic(&mw.BadRequestError{Message: constants.IdNotInt})
	}

	comment, err := h.NewsUseCase.GetCommentById(commentIdInt)

	if err != nil {
		panic(err)
	}

	response.Success(w, http.StatusOK, comment)
}
