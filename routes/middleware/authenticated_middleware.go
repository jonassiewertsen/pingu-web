package middleware

import (
	"github.com/gofiber/fiber/v2"
	"pingu-web/application"
)

func Authenticated(app *application.App) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Check if the user is authenticated in the session
		s, _ := app.Session.Get(c)
		userId := s.Get("userId")
		if userId == nil {
			// Redirect to login if not authenticated
			return c.Redirect("/login")
		}

		// Return err if exists, else move to next handler
		return c.Next()
	}
}
