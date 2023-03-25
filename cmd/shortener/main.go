package main

import (
	"fmt"
	"log"

	"github.com/casnerano/go-url-shortener/internal/app/server"
	"github.com/casnerano/go-url-shortener/internal/app/service/cleaner"
)

// These variables are configured using ldflags.
//
// For example:
//  go run -ldflags "-X main.Version=v1.0.1 \
//  -X 'main.buildVersion=1.0.0'" ./cmd/shortener/main.go
var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {

	fmt.Printf(
		"Build version: %s\nBuild date: %s\nBuild commit: %s\n",
		buildVersion,
		buildDate,
		buildCommit,
	)

	app := server.NewApplication()
	defer app.Shutdown()

	if ttl := app.Config.ShortURL.TTL; ttl > 0 {
		go cleaner.New(app.Store).CleanOlderShortURL(ttl)
	}

	log.Fatal(app.RunServer())
}
