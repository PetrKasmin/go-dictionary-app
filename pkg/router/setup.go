package router

import (
	"github.com/app-dictionary/app/controller"
	"github.com/app-dictionary/app/middleware"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func InstallRouter(app *fiber.App, embedFS http.FileSystem) {
	setup(
		app,
		controller.NewAppController(embedFS),
		middleware.NewMiddleware(),
		NewApiRouter(),
		NewHttpRouter(),
	)
}

func setup(
	app *fiber.App,
	ctr controller.Controller,
	mdw middleware.Middleware,
	routers ...Router,
) {
	for _, r := range routers {
		r.InstallRouter(app, ctr, mdw)
	}
}
