package routes

import (
	"pingu-web/application"
	"pingu-web/controller"
)

func initWebRoutes(app *application.App) {
	ctl := controller.NewController(app)

	app.Fiber.Get("/", ctl.Home)
	// Add more routes
}
