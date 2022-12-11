package main

import (
	"log"

	"github.com/casnerano/go-url-shortener/internal/app/server"
	"github.com/casnerano/go-url-shortener/internal/app/service/cleaner"
)

func main() {
	app := server.NewApplication()
	defer app.CloseResources()

	if ttl := app.Config.ShortURL.TTL; ttl > 0 {
		go cleaner.New(app.Store).CleanOlderShortURL(ttl)
	}

	log.Fatal(app.RunServer())
}
