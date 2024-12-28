package routes

import (
	"github.com/gofiber/fiber/v2"
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

	// 404 Handler, in case a route could not be found
	app.Fiber.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("errors/404", fiber.Map{})
	})
}
