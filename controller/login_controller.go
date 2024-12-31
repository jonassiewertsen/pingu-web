package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetLogin(ctx *fiber.Ctx) error {
	s, _ := c.App.Session.Get(ctx)

	userId := s.Get("userId")
	if userId != nil {
		c.App.Log.Info("Already authenticated. User ID: " + userId.(string))
	} else {
		c.App.Log.Info("User not authenticated. Authenticate")

		s.Set("userId", "1234")
		_ = s.Save()
	}

	c.render(ctx, "index", "layouts/main", fiber.Map{
		"title": "Login",
	})

	return nil
}

func (c *Controller) GetLogout(ctx *fiber.Ctx) error {
	s, _ := c.App.Session.Get(ctx)
	_ = s.Destroy()

	_ = ctx.RedirectToRoute("home", nil)

	return nil
}

func (c *Controller) GetInternal(ctx *fiber.Ctx) error {
	c.render(ctx, "index", "layouts/main", fiber.Map{
		"title": "Internal area",
	})

	return nil
}
