package routes

import (
	"pingu-web/application"
	"pingu-web/routes/middleware"
)

/**
 * Initialize Middleware
 *
 * Check out Fiber examples:
 * @see https://docs.gofiber.io/next/middleware/csrf
 */
func initMiddleware(app *application.App) {
	app.Fiber.Use(middleware.RecoverMiddleware())

	// app.Fiber.Use(func(c *fiber.Ctx) error {
	// 	c.Set("X-Request-ID", "123")
	// 	return c.Next()
	// })
}
