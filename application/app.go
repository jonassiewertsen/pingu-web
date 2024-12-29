package application

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log/slog"
)

type App struct {
	Cache   fiber.Storage
	config  *Config
	Fiber   *fiber.App
	Log     *slog.Logger
	Session *session.Store
}

func Create(config *Config, viewPath embed.FS) *App {
	templateEngine := initializeTemplateEngine(viewPath)
	cacheStorage := initializeCache(config)
	sessionStorage := initializeSessions(config)
	logger := initializeLogging(config)

	// Pass the template engine to the Views
	f := fiber.New(fiber.Config{
		Views: templateEngine,
	})

	return &App{
		Cache:   cacheStorage,
		config:  config,
		Fiber:   f,
		Log:     logger,
		Session: sessionStorage,
	}
}

func (a *App) GetConfig() *Config {
	return a.config
}
