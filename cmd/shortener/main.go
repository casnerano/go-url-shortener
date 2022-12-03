package main

import (
	"log"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/server"
	"github.com/casnerano/go-url-shortener/internal/app/service/cleaner"
)

func main() {
	app := server.NewApplication()

	if app.Config.Storage.Type == config.StorageTypeDatabase {
		if err := app.LoadMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	if ttl := app.Config.ShortURL.TTL; ttl > 0 {
		go cleaner.New(app.Store).CleanOlderShortURL(ttl)
	}

	log.Fatal(app.RunServer())
}
