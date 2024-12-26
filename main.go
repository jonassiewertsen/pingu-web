package main

import (
	"embed"
	"fmt"
	"pingu-web/application"
	"pingu-web/routes"
)

//go:embed resources/views
var viewsPath embed.FS

func main() {
	fmt.Printf("JO")

	cfg := application.NewConfig()
	app := application.Create(cfg, viewsPath)

	routes.Init(app)

	// Serve static assets from the public folder
	app.Fiber.Static("/", "./public")

	err := app.Fiber.Listen(cfg.AppPort)
	if err != nil {
		panic(err) // Or any other error handling you prefer
	}
}
