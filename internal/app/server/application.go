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

type application struct {
	Config *config.Config
	Store  repository.Store
	router *chi.Mux
}

func NewApplication() *application {
	app := &application{}
	app.init()
	return app
}

func (app *application) init() {
	app.initConfig()
	app.initRepositoryStore()
	app.initRouter()
	app.initRoutes()
}

// Запуск сервера
func (app *application) RunServer() error {
	return http.ListenAndServe(app.Config.ServerAddr, app.router)
}

// Путь к файлу конфигурации из параметров запуска приложения
func (app *application) extractAppConfigName() string {
	confName := flag.String("config", "", "app configuration filename")
	flag.Parse()
	return *confName
}

// Инициализация конфигурации
func (app *application) initConfig() {
	app.Config = config.New()

	if cfgName := app.extractAppConfigName(); cfgName != "" {
		if err := config.Unmarshal(cfgName, app.Config); err != nil {
			log.Printf("failed to read file %s", cfgName)
		}
	}
}

func (app *application) initRouter() {
	app.router = chi.NewRouter()
}

// Группа репозиториев для хранилища
// По умолчанию - хранилище в озу - memstore.Store
func (app *application) initRepositoryStore() {
	switch app.Config.Storage.Type {
	default:
		app.Store = memstore.NewStore()
	}
}

// Инициализация роутов
func (app *application) initRoutes() {
	shortener := app.getShortenerHandlerGroup()

	app.router.Get("/{shortCode}", shortener.URLGetHandler)
	app.router.Post("/", shortener.URLPostHandler)
}

// Сервис для сокращения URL
func (app *application) getURLHashService() (h hash.Hash) {
	h, _ = hash.NewRandom(5, 10)
	return
}

// Группа обработчиков для сокращения URL
func (app *application) getShortenerHandlerGroup() *handler.Shortener {
	URLRepository := app.Store.URL()
	hashService := app.getURLHashService()
	return handler.NewShortener(app.Config, URLRepository, hashService)
}

func (app *application) LoadMigrations() error {
	migrator := migration.NewManager(app.Config)
	return migrator.LoadMigrations()
}
