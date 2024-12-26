package application

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"net/http"
)

type App struct {
	Fiber  *fiber.App
	config *Config
}

func Create(config *Config, viewPath embed.FS) *App {
	config.Fiber.Views = django.NewPathForwardingFileSystem(http.FS(viewPath), "/resources/views", ".django.html")

	// Pass the engine to the Views
	f := fiber.New(fiber.Config{
		Views: config.Fiber.Views,
	})

	return &App{
		Fiber:  f,
		config: config,
	}
}

func (a *App) GetConfig() *Config {
	return a.config
}
