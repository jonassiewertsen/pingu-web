package routes

import (
	"github.com/gofiber/fiber/v2"
	"pingu-web/application"
)

func Init(app *application.App) {
	// TODO: Add a controller
	app.Fiber.Get("/", func(ctx *fiber.Ctx) error {
		// Render index
		return ctx.Render("index", fiber.Map{
			"title":      "Hello, World!",
			"stuff_list": []string{"a", "b", "c"},
			"ViteDev":    app.GetConfig().AppIsDev,
		}, "layouts/main")
	})
}
