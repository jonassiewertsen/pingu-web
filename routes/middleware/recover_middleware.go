package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"runtime/debug"
)

/**
 * RecoverMiddleware recovers from panics and returns a 500 error page instead
 *
 * Fiber has a built-in recover middleware, but the Fiber middleware prints the panic message.
 * For procution use, I don't want to show the panic message to the user, as it might
 * contain sensitive information. Therefore, a custom recover middleware is created.
 */
func RecoverMiddleware() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) (err error) {

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				// Log the full panic message and stack trace
				fmt.Printf("Panic recovered: %v\n%s\n", r, debug.Stack())

				// Optionally, log to a file or external logging system

				// Return 500 status and error message
				_ = c.Status(fiber.StatusInternalServerError).Render("errors/500", nil)
			}
		}()

		// Return err if exists, else move to next handler
		return c.Next()
	}
}
