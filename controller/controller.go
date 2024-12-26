package controller

import (
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
