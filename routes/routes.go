package routes

import (
	"pingu-web/application"
)

func Init(app *application.App) {
	initWebRoutes(app)
	// initApiRoutes(app)
}
