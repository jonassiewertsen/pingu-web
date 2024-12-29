package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Home(ctx *fiber.Ctx) error {
	c.App.Log.Info("Home page accessed")

	c.render(ctx, "index", "layouts/main", fiber.Map{
		"title":      "Hello, World!",
		"stuff_list": []string{"a", "b", "c"},
	})

	return nil
}
