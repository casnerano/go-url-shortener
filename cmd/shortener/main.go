package main

import (
	"log"
	"net/http"

	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func main() {
	mux := http.NewServeMux()
	shortURLRepository := repository.NewShortURL(storage.NewInMemory())
	randHashService, _ := hash.NewRandom(5, 10)
	shortener := handler.NewShortener(shortURLRepository, randHashService)

	mux.HandleFunc("/", shortener.URLHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
