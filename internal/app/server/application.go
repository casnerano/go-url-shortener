package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/repository/sqlstore"
	"github.com/casnerano/go-url-shortener/internal/app/service/hash"
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

// Запуск сервера
func (app *Application) RunServer() error {
	fmt.Printf("Server started: %s\n", app.Config.ServerAddr)
	fmt.Printf("Use storage is %s\n", app.Config.Storage.Type)
	return http.ListenAndServe(app.Config.ServerAddr, app.router)
}

// Путь к файлу конфигурации из параметров запуска приложения
func (app *Application) extractAppConfigName() string {
	confName := flag.String("config", "", "app configuration filename")
	flag.Parse()
	return *confName
}

// Инициализация конфигурации
func (app *Application) initConfig() {
	app.Config = config.New()

	if cfgName := app.extractAppConfigName(); cfgName != "" {
		if err := config.Unmarshal(cfgName, app.Config); err != nil {
			log.Fatalf("failed to read file %s", cfgName)
		}
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
	default:
		app.Store = memstore.NewStore()
	}
}

// Инициализация роутов
func (app *Application) initRoutes() {
	shortener := app.getShortenerHandlerGroup()

	app.router.Get("/{shortCode}", shortener.URLGetHandler)
	app.router.Post("/", shortener.URLPostHandler)
}

// Сервис для сокращения URL
func (app *Application) getURLHashService() (h hash.Hash) {
	h, _ = hash.NewRandom(5, 10)
	return
}

// Группа обработчиков для сокращения URL
func (app *Application) getShortenerHandlerGroup() *handler.Shortener {
	URLRepository := app.Store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortener(app.Config, URLRepository, hashService)
}
