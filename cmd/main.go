package main

import (
	"hackernews-service/internal/delivery"
	"hackernews-service/internal/usecase"
	"log"
	"net/http"
)

func main() {

	newsUsecase := usecase.NewNewsUsecase()

	newsHandler := &delivery.NewsHandler{
		NewsUseCase: newsUsecase,
	}

	router := delivery.NewRouter(newsHandler)

	
	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}