package application

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/django/v3"
	"net/http"
)

type App struct {
	Cache   fiber.Storage
	Fiber   *fiber.App
	Session *session.Store
	config  *Config
}

func Create(config *Config, viewPath embed.FS) *App {
	config.Fiber.Views = django.NewPathForwardingFileSystem(http.FS(viewPath), "/resources/views", ".django.html")

	// Pass the engine to the Views
	f := fiber.New(fiber.Config{
		Views: config.Fiber.Views,
	})

	cacheStorage := createCacheStorage(config)
	sessionStorage := createSessionStore(config)

	return &App{
		Fiber:   f,
		Cache:   cacheStorage,
		Session: sessionStorage,
		config:  config,
	}
}

func (a *App) GetConfig() *Config {
	return a.config
}
