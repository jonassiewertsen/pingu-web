package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"net/http"
)

//go:embed resources/views
var viewsAsssets embed.FS

func main() {
	// Detect the environment (default to production)
	// env := os.Getenv("APP_ENV") // GET from .env later
	isDev := true

	// Create a new engine
	engine := django.NewPathForwardingFileSystem(http.FS(viewsAsssets), "/resources/views", ".django.html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware to set development mode
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("ViteDev", isDev)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"title":      "Hello, World dd!",
			"stuff_list": []string{"a", "b", "c"},
			"ViteDev":    c.Locals("ViteDev"),
		}, "layouts/main")
	})

	// Serve static assets in production
	app.Static("/", "./public")

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
