package controller

import (
	"github.com/gofiber/fiber/v2"
	"pingu-web/application"
)

type Controller struct {
	App *application.App
}

func NewController(app *application.App) *Controller {
	return &Controller{
		App: app,
	}
}

/**
 * Used to render a view. On top of simply rendering the view, we also:
 * - Pass the AppIsDev config to the view
 */
func (c *Controller) render(ctx *fiber.Ctx, template string, layout string, data fiber.Map) {
	data["ViteDev"] = c.App.GetConfig().AppIsDev

	err := ctx.Render(template, data, layout)

	// In case an error occurred, we will log it and show a 500 error page
	if err != nil {
		// TODO: Log error
		_ = ctx.Render("errors/500", fiber.Map{})
	}
}
