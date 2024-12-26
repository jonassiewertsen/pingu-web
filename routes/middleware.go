package routes

import "pingu-web/application"

/**
 * Initialize Middleware
 *
 * Check out Fiber examples:
 * @see https://docs.gofiber.io/next/middleware/csrf
 */
func initMiddleware(app *application.App) {
	// app.Fiber.Use(func(c *fiber.Ctx) error {
	// 	c.Set("X-Request-ID", "123")
	// 	return c.Next()
	// })
}
