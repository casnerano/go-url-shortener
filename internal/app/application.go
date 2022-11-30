package app

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/handler"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
)

type Application struct {
	cfg    *config.Config
	router *chi.Mux
}

func NewApplication() *Application {
	return &Application{
		router: chi.NewRouter(),
	}
}

func (app *Application) Run() error {
	app.initConfig()
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
	app.cfg = config.New()

	if cfgName := app.extractAppConfigName(); cfgName != "" {
		_ = config.Unmarshal(cfgName, app.cfg)
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
	return http.ListenAndServe(app.cfg.ServerAddr, app.router)
}

// Текущий репозиторий для URL
// По умолчанию - repository.Memory
func (app *Application) getURLRepository() (rep repository.URLRepository) {
	switch app.cfg.Storage.Type {
	default:
		rep = repository.NewMemory()
	}
	return
}

// Сервис для сокращения URL
func (app *Application) getURLHashService() (h hash.Hash) {
	h, _ = hash.NewRandom(5, 10)
	return
}

// Группа обработчиков для сокращения URL
func (app *Application) getShortenerHandlerGroup() *handler.Shortener {
	URLRepository := app.getURLRepository()
	hashService := app.getURLHashService()
	return handler.NewShortener(app.cfg, URLRepository, hashService)
}
