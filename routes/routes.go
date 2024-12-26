package routes

import (
	"pingu-web/application"
)

/**
 * Initialize routes
 *
 * To keep the code clean, we will initialize the routes in separate files.
 * This is only a suggestion, feel free doing it your way if your way of course.
 */
func Init(app *application.App) {
	initMiddleware(app)
	initWebRoutes(app)
	// initApiRoutes(app)
}
