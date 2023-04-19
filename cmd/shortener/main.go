package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"github.com/casnerano/go-url-shortener/internal/app/server"
	"github.com/casnerano/go-url-shortener/internal/app/service/cleaner"
)

// These variables are configured using ldflags.
//
// For example:
// go run -ldflags "-X main.Version=v1.0.1 \
// -X 'main.buildVersion=1.0.0'" ./cmd/shortener/main.go
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

	wg := &sync.WaitGroup{}
	app := server.NewApplication()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if ttl := app.GetConfig().ShortURL.TTL; ttl > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleaner.New(app.GetStore()).CleanOlderShortURL(ctx, ttl)
		}()
	}

	go func() {
		if err := app.RunHTTPServer(); err != nil && err != http.ErrServerClosed {
			log.Fatal(
				fmt.Sprintf("Failed to start http server at %s", app.GetConfig().Server.Addr),
				err,
			)
		}
	}()

	go func() {
		if err := app.RunGRPCServer(); err != nil {
			log.Fatal(
				fmt.Sprintf("Failed to start grpc server at %s", app.GetConfig().Server.Addr),
				err,
			)
		}
	}()

	<-ctx.Done()

	wg.Wait()

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
