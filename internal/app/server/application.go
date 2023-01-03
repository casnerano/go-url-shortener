package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/middleware"
	"github.com/casnerano/go-url-shortener/internal/app/migration"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/filestore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/sqlstore"
	"github.com/casnerano/go-url-shortener/internal/app/service"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
)

type Application struct {
	Config  *config.Config
	Store   repository.Store
	router  *chi.Mux
	pgxpool *pgxpool.Pool
}

func NewApplication() *Application {
	app := &Application{}
	app.init()
	return app
}

func (app *Application) init() {
	app.initConfig()
	if app.Config.Storage.Type == config.StorageTypeDatabase {
		// todo: logging
		_ = app.loadDBMigrations(app.Config.Storage.DSN)
	}
	app.initRepositoryStore()
	app.initRouter()
	app.initRoutes()
}

func (app *Application) Shutdown() error {
	if closer, ok := app.Store.(io.Closer); ok {
		_ = closer.Close()
	}
	if app.pgxpool != nil {
		app.pgxpool.Close()
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
		log.Fatal(err.Error())
		// todo: logging
	}

	if err := app.Config.SetEnvironmentValues(); err != nil {
		log.Fatal(err.Error())
		// todo: logging
	}

	if err := app.Config.SetAppFlagValues(); err != nil {
		log.Fatal(err.Error())
		// todo: logging
	}

	if app.Config.Storage.DSN != "" {
		app.Config.Storage.Type = config.StorageTypeDatabase
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

// Инициализация роутов
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
		r.Get("/user/urls", shortURL.GetUserURLHistory)
	})

	app.router.Get("/ping", database.PingPostreSQL)
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

// Группа обработчиков для работы с базой данных
func (app *Application) getDatabaseHandlerGroup() *handler.Database {
	_pgxpool, err := app.getDBConnection()
	if err != nil {
		panic(err)
	}
	return handler.NewDatabase(_pgxpool)
}

// Подключение к базе данных
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

// Загрузка миграций базы данных
func (app *Application) loadDBMigrations(dbDSN string) error {
	migrator := migration.NewMigrator(migration.MigrationDBTypePostgres, dbDSN)
	return migrator.LoadMigrations()
}
