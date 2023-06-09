package delivery

import (
	"net/http"

	mw "hackernews-service/internal/delivery/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(newsHandler *NewsHandler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	router.Use(mw.ErrorHandler)


	router.Use(middleware.Heartbeat("/ping"))

	router.Route("/api", func(router chi.Router) {
		router.Get("/news", newsHandler.GetAll)
		router.Get("/news/story/{storyId}", newsHandler.GetStoryById)
		router.Get("/news/comment/{commentId}", newsHandler.GetCommentById)
	})

	return router

}