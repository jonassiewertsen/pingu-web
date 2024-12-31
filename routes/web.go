package routes

import (
	"pingu-web/application"
	"pingu-web/controller"
	"pingu-web/routes/middleware"
)

func initWebRoutes(app *application.App) {
	ctl := controller.NewController(app)

	app.Fiber.Get("/", ctl.Home).Name("home")
	app.Fiber.Get("/login", ctl.GetLogin).Name("login")
	app.Fiber.Get("/logout", ctl.GetLogout).Name("logout")

	// Authenticated routes
	internal := app.Fiber.Group("/internal", middleware.Authenticated(app))
	internal.Get("/", middleware.Authenticated(app), ctl.GetInternal).Name("internal")
	// Add more routes
}
