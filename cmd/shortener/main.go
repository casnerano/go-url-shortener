package main

import (
	"log"

	"github.com/casnerano/go-url-shortener/internal/app/server"
)

func main() {
	app := server.NewApplication()
	log.Fatal(app.Run())
}
