package main

import (
	"hackernews-service/internal/delivery"
	firebaserepository "hackernews-service/internal/repository/firebase"
	"hackernews-service/internal/usecase"
	memorycache "hackernews-service/pkg/memory_cache"
	"log"
	"net/http"
)

func main() {

	newsFirebaseRepository := firebaserepository.NewNewsFirebaseRepository(http.DefaultClient)
	cache := memorycache.NewMemoryCache()
	newsUsecase := usecase.NewNewsUsecase(newsFirebaseRepository, cache)
	newsUsecase.StartCacheUpdate()

	newsHandler := &delivery.NewsHandler{
		NewsUseCase: newsUsecase,
	}

	router := delivery.NewRouter(newsHandler)

	
	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}