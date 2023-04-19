package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/acme/autocert"
	google_grpc "google.golang.org/grpc"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/filestore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/sqlstore"
	"github.com/casnerano/go-url-shortener/internal/app/server/grpc"
	pb "github.com/casnerano/go-url-shortener/internal/app/server/grpc/proto"
	"github.com/casnerano/go-url-shortener/internal/app/server/http/handler"
	"github.com/casnerano/go-url-shortener/internal/app/server/http/middleware"
	"github.com/casnerano/go-url-shortener/internal/app/service"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
)

// Application the structure that is responsible for all dependencies
// and contains the methods of launching the application.
type Application struct {
	httpServer *http.Server
	grpServer  *google_grpc.Server
	config     *config.Config
	store      repository.Store
	router     *chi.Mux
	pgxpool    *pgxpool.Pool
}

// NewApplication - constructor.
func NewApplication() *Application {
	app := &Application{
		router: chi.NewRouter(),
	}
	app.init()
	return app
}

func (app *Application) init() {
	app.initConfig()
	app.initRepositoryStore()
}

// GetStore getter for app store
func (app *Application) GetStore() repository.Store {
	return app.store
}

// GetConfig getter for app config
func (app *Application) GetConfig() *config.Config {
	return app.config
}

// Shutdown closes all open resources.
func (app *Application) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	if closer, ok := app.store.(io.Closer); ok {
		_ = closer.Close()
	}
	if app.pgxpool != nil {
		app.pgxpool.Close()
	}
	return nil
}

// RunHTTPServer run HTTP server.
func (app *Application) RunHTTPServer() error {
	app.initHTTPRoutes()

	fmt.Printf("HTTP Server started: %s\n", app.config.Server.Addr)

	app.httpServer = &http.Server{
		Addr:    app.config.Server.Addr,
		Handler: app.router,
	}

	if app.config.Server.EnableHTTPS {
		autoCertManager := &autocert.Manager{
			Cache:      autocert.DirCache("./var"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("shortener.ru", "www.shortener.ru"),
		}

		app.httpServer.TLSConfig = autoCertManager.TLSConfig()
		return app.httpServer.ListenAndServeTLS("", "")
	}

	return app.httpServer.ListenAndServe()
}

// RunGRPCServer run GRPC server.
func (app *Application) RunGRPCServer() error {
	fmt.Println("GRPC Server started: :3200")

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}

	app.grpServer = google_grpc.NewServer()
	pb.RegisterShortenerServer(app.grpServer, &grpc.ShortenerServer{})

	return app.grpServer.Serve(listen)
}

// Initialization configs.
func (app *Application) initConfig() {
	app.config = config.New()
	app.config.Init()
}

func (app *Application) initRepositoryStore() {
	switch app.config.Storage.Type {
	case config.StorageTypeDatabase:
		_pgxpool, err := app.getDBConnection()
		if err != nil {
			panic(err)
		}
		app.store = sqlstore.NewStore(_pgxpool)
	case config.StorageTypeFile:
		store, err := filestore.NewStore(app.config.Storage.Path)
		if err != nil {
			panic(err)
		}
		app.store = store
	default:
		app.store = memstore.NewStore()
	}
}

// Initialization routes.
func (app *Application) initHTTPRoutes() {
	shortURL := app.getShortURLHandlerGroup()
	database := app.getDatabaseHandlerGroup()

	app.router.Use(middleware.Authenticate([]byte(app.config.App.Secret)))
	app.router.Use(middleware.GzipCompress(1400))
	app.router.Use(middleware.GzipDecompress())

	app.router.Get("/{shortCode}", shortURL.GetOriginalURL)
	app.router.Post("/", shortURL.PostText)

	app.router.Route("/api", func(r chi.Router) {
		r.Route("/internal/stats", func(r chi.Router) {
			if app.config.Server.TrustedSubnet != "" {
				r.Use(middleware.TrustedSubnet(app.config.Server.TrustedSubnet))
			}
			r.Get("/", shortURL.GetStats)
		})

		r.Post("/shorten", shortURL.PostJSON)
		r.Post("/shorten/batch", shortURL.PostBatchJSON)
		r.Get("/user/urls", shortURL.GetUserURLHistory)
		r.Delete("/user/urls", shortURL.DeleteBatchJSON)
	})

	app.router.Get("/ping", database.PingPostreSQL)
}

// Service for url shortener.
func (app *Application) getURLHashService() (h hasher.Hash) {
	h = hasher.NewUnique()
	return
}

// Handler group for url shortener.
func (app *Application) getShortURLHandlerGroup() *handler.ShortURL {
	URLRepository := app.store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortURL(app.config, service.NewURL(URLRepository, hashService))
}

// Handler group for database.
func (app *Application) getDatabaseHandlerGroup() *handler.Database {
	_pgxpool, err := app.getDBConnection()
	if err != nil {
		panic(err)
	}
	return handler.NewDatabase(_pgxpool)
}

// Database connection.
func (app *Application) getDBConnection() (*pgxpool.Pool, error) {
	if app.pgxpool != nil {
		return app.pgxpool, nil
	}

	dsn := app.config.Storage.DSN

	var err error
	app.pgxpool, err = pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return app.pgxpool, nil
}
