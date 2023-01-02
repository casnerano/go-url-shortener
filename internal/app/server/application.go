package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

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

type Application struct {
	Config *config.Config
	Store  repository.Store
	router *chi.Mux
}

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

func (app *Application) Shutdown() error {
	if closer, ok := app.Store.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// Запуск сервера
func (app *Application) RunServer() error {
	fmt.Printf("Server started: %s\n", app.Config.Server.Addr)
	fmt.Printf("Use storage is %s\n", app.Config.Storage.Type)
	return http.ListenAndServe(app.Config.Server.Addr, app.router)
}

// Инициализация конфигурации
func (app *Application) initConfig() {
	app.Config = config.New()
	app.Config.SetDefaultValues()

	if err := app.Config.SetConfigFileValues(); err != nil {
		// todo: logging
	}

	if err := app.Config.SetEnvironmentValues(); err != nil {
		// todo: logging
	}

	if err := app.Config.SetAppFlagValues(); err != nil {
		// todo: logging
	}

	if app.Config.Storage.Path != "" {
		app.Config.Storage.Type = config.StorageTypeFile
	}
}

func (app *Application) initRouter() {
	app.router = chi.NewRouter()
}

// Группа репозиториев для хранилища
// По умолчанию - хранилище в озу - memstore.Store
func (app *Application) initRepositoryStore() {
	switch app.Config.Storage.Type {
	case config.StorageTypeDatabase:
		dsn := app.Config.Storage.DSN
		pgxConn, err := pgx.Connect(context.TODO(), dsn)
		if err != nil {
			log.Fatalf("Failed to connect to the database using dsn \"%s\"", dsn)
		}
		app.Store = sqlstore.NewStore(pgxConn)
	case config.StorageTypeFile:
		store, err := filestore.NewStore(app.Config.Storage.Path)
		if err != nil {
			log.Fatalf("Failed to initialization file-storage: \"%s\"", err)
		}
		app.Store = store
	default:
		app.Store = memstore.NewStore()
	}
}

// Инициализация роутов
func (app *Application) initRoutes() {
	shortURL := app.getShortURLHandlerGroup()

	app.router.Use(middleware.Authenticate([]byte("#easy_secret_key")))
	app.router.Use(middleware.GzipCompress(1400))
	app.router.Use(middleware.GzipDecompress())

	app.router.Get("/{shortCode}", shortURL.GetOriginalURL)
	app.router.Post("/", shortURL.PostText)

	app.router.Route("/api", func(r chi.Router) {
		r.Post("/shorten", shortURL.PostJSON)
		r.Get("/user/urls", shortURL.GetUserURLHistory)
	})
}

// Сервис для сокращения URL
func (app *Application) getURLHashService() (h hasher.Hash) {
	h, _ = hasher.NewRandom(5, 10)
	return
}

// Группа обработчиков для сокращения URL
func (app *Application) getShortURLHandlerGroup() *handler.ShortURL {
	URLRepository := app.Store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortURL(app.Config, service.NewURL(URLRepository, hashService))
}
