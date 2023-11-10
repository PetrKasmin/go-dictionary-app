package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func InstallRouter(app *fiber.App, embedFS http.FileSystem) {
	setup(app, NewApiRouter(), NewHttpRouter(embedFS))
}

func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}
