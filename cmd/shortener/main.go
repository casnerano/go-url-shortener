package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func main() {
	// Дефолтная конфигурация
	conf := config.New()

	// Пытаемся подключить конфиг. файл из переменных окружения
	if configFilename := os.Getenv("CONFIG_FILENAME"); configFilename != "" {
		_ = config.Unmarshal(configFilename, conf)
	}

	store := storage.NewMemory()
	if conf.Storage.Type == config.STORAGE_TYPE_MEMORY {
	}

	shortURLRepository := repository.NewShortURL(store)
	randHashService, _ := hash.NewRandom(5, 10)
	shortener := handler.NewShortener(shortURLRepository, randHashService)

	router := chi.NewRouter()

	router.Get("/{shortCode}", shortener.URLGetHandler)
	router.Post("/", shortener.URLPostHandler)

	log.Fatal(http.ListenAndServe(conf.ServerAddr, router))
}
