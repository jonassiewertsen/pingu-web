package application

import (
	"embed"
	"github.com/gofiber/template/django/v3"
	"net/http"
)

func initializeTemplateEngine(viewPath embed.FS) *django.Engine {
	return django.NewPathForwardingFileSystem(http.FS(viewPath), "/resources/views", ".django.html")
}
