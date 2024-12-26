package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Home(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"title":      "Hello, World!",
		"stuff_list": []string{"a", "b", "c"},
		"ViteDev":    c.App.GetConfig().AppIsDev,
	}, "layouts/main")
}
