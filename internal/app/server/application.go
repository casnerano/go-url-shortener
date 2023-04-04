package server

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/acme/autocert"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/middleware"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/filestore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/sqlstore"
	"github.com/casnerano/go-url-shortener/internal/app/service"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
)

// Application the structure that is responsible for all dependencies
// and contains the methods of launching the application.
type Application struct {
	Config  *config.Config
	Store   repository.Store
	router  *chi.Mux
	pgxpool *pgxpool.Pool
}

// NewApplication - constructor.
func NewApplication() *Application {
	app := &Application{}
	app.init()
	return app
}

func (app *Application) init() {
	app.initConfig()
	app.initRepositoryStore()
	app.initRouter()
	app.initRoutes()
}

// Shutdown closes all open resources.
func (app *Application) Shutdown() error {
	if closer, ok := app.Store.(io.Closer); ok {
		_ = closer.Close()
	}
	if app.pgxpool != nil {
		app.pgxpool.Close()
	}
	return nil
}

// RunServer run server.
func (app *Application) RunServer() error {
	fmt.Printf("Server started: %s\n", app.Config.Server.Addr)
	fmt.Printf("Use storage is %s\n", app.Config.Storage.Type)

	srv := &http.Server{
		Addr:    app.Config.Server.Addr,
		Handler: app.router,
	}

	if app.Config.Server.EnableHTTPS {
		autoCertManager := &autocert.Manager{
			Cache:      autocert.DirCache("./var"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("shortener.ru", "www.shortener.ru"),
		}

		srv.TLSConfig = autoCertManager.TLSConfig()
		return srv.ListenAndServeTLS("", "")
	}

	return srv.ListenAndServe()
}

// Initialization configs.
func (app *Application) initConfig() {
	app.Config = config.New()
	app.Config.Init()
}

func (app *Application) initRouter() {
	app.router = chi.NewRouter()
}

func (app *Application) initRepositoryStore() {
	switch app.Config.Storage.Type {
	case config.StorageTypeDatabase:
		_pgxpool, err := app.getDBConnection()
		if err != nil {
			panic(err)
		}
		app.Store = sqlstore.NewStore(_pgxpool)
	case config.StorageTypeFile:
		store, err := filestore.NewStore(app.Config.Storage.Path)
		if err != nil {
			panic(err)
		}
		app.Store = store
	default:
		app.Store = memstore.NewStore()
	}
}

// Initialization routes.
func (app *Application) initRoutes() {
	shortURL := app.getShortURLHandlerGroup()
	database := app.getDatabaseHandlerGroup()

	app.router.Use(middleware.Authenticate([]byte(app.Config.App.Secret)))
	app.router.Use(middleware.GzipCompress(1400))
	app.router.Use(middleware.GzipDecompress())

	app.router.Get("/{shortCode}", shortURL.GetOriginalURL)
	app.router.Post("/", shortURL.PostText)

	app.router.Route("/api", func(r chi.Router) {
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
	URLRepository := app.Store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortURL(app.Config, service.NewURL(URLRepository, hashService))
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

	dsn := app.Config.Storage.DSN

	var err error
	app.pgxpool, err = pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return app.pgxpool, nil
}
