package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/migration"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
)

type Application struct {
	Config *config.Config
	store  repository.Store
	router *chi.Mux
}

func NewApplication() *Application {
	return &Application{
		router: chi.NewRouter(),
	}
}

func (app *Application) Run() error {
	app.initConfig()
	app.initRepositoryStore()
	app.initRoutes()
	return app.runServer()
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
			log.Printf("failed to read file %s", cfgName)
		}
	}
}

// Группа репозиториев для хранилища
// По умолчанию - хранилище в озу - memstore.Store
func (app *Application) initRepositoryStore() {
	switch app.Config.Storage.Type {
	default:
		app.store = memstore.NewStore()
	}
}

// Инициализация роутов
func (app *Application) initRoutes() {
	shortener := app.getShortenerHandlerGroup()

	app.router.Get("/{shortCode}", shortener.URLGetHandler)
	app.router.Post("/", shortener.URLPostHandler)
}

// Запуск сервера
func (app *Application) runServer() error {
	return http.ListenAndServe(app.Config.ServerAddr, app.router)
}

// Сервис для сокращения URL
func (app *Application) getURLHashService() (h hash.Hash) {
	h, _ = hash.NewRandom(5, 10)
	return
}

// Группа обработчиков для сокращения URL
func (app *Application) getShortenerHandlerGroup() *handler.Shortener {
	URLRepository := app.store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortener(app.Config, URLRepository, hashService)
}

func (app *Application) LoadMigrations() error {
	migrator := migration.NewManager(app.Config)
	return migrator.LoadMigrations()
}
