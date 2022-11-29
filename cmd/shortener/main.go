package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func main() {
	shortURLRepository := repository.NewShortURL(storage.NewInMemory())
	randHashService, _ := hash.NewRandom(5, 10)
	shortener := handler.NewShortener(shortURLRepository, randHashService)

	router := chi.NewRouter()

	router.Get("/{shortCode}", shortener.URLGetHandler)
	router.Post("/", shortener.URLPostHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
