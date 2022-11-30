package main

import (
	"log"

	"github.com/casnerano/go-url-shortener/internal/app"
)

func main() {
	application := app.NewApplication()
	log.Fatal(application.Run())
}
