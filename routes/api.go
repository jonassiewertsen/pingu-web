package routes

import (
	"pingu-web/application"
	"pingu-web/controller"
)

var apiPrefix = "/api/v1"

func initApiRoutes(app *application.App) {
	ctl := controller.NewController(app)

	app.Fiber.Get(apiPrefix+"/", ctl.Home)
	// Add more routes
}
