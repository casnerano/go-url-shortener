package main

import (
	"log"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/server"
)

func main() {
	app := server.NewApplication()

	if app.Config.Storage.Type == config.StorageTypeDatabase {
		if err := app.LoadMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	log.Fatal(app.Run())
}
