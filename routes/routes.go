package routes

import (
	"pingu-web/application"
)

func Init(app *application.App) {
	/**
	 * Add Middleware
	 *
	 * Check out Fiber examples:
	 * @see https://docs.gofiber.io/next/middleware/csrf
	 */
	// app.Fiber.Use(csrf.New())

	/**
	 * Initialize routes
	 *
	 * To keep the code clean, we will initialize the routes in separate files.
	 * This is only a suggestion, feel free doing it your way if your way of course.
	 */
	initWebRoutes(app)
	// initApiRoutes(app)
}
