package main

import (
	"embed"
	"pingu-web/application"
	"pingu-web/routes"
)

//go:embed resources/views
var viewsPath embed.FS

func main() {
	// Initialize the configuration
	cfg := application.NewConfig()
	
	// Initialize the application itself
	app := application.Create(cfg, viewsPath)

	// Serve static assets from the public folder
	app.Fiber.Static("/", "./public")

	// Initialize routes, which should happen after the static assets
	// This way we can catch 404 errors.
	// @see routes/routes.go
	routes.Init(app)

	err := app.Fiber.Listen(cfg.AppPort)
	if err != nil {
		panic(err) // Or any other error handling you prefer
	}
}
