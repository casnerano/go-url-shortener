package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
)

func main() {
	// Дефолтная конфигурация
	conf := config.New()

	// Пытаемся подключить конфиг. файл из параметров запуска приложения
	confName := flag.String("config", "", "app configuration filename")
	flag.Parse()

	if *confName != "" {
		_ = config.Unmarshal(*confName, conf)
	}

	var URLRepository repository.URLRepository
	switch conf.Storage.Type {
	default:
		URLRepository = repository.NewMemory()
	}

	randHashService, _ := hash.NewRandom(5, 10)
	shortener := handler.NewShortener(conf, URLRepository, randHashService)

	router := chi.NewRouter()

	router.Get("/{shortCode}", shortener.URLGetHandler)
	router.Post("/", shortener.URLPostHandler)

	log.Fatal(http.ListenAndServe(conf.ServerAddr, router))
}
