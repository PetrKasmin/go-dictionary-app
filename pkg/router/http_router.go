package router

import (
	"github.com/app-dictionary/app/controller"
	"github.com/app-dictionary/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type HttpRouter struct {
}

func (h *HttpRouter) InstallRouter(
	app *fiber.App,
	ctr controller.Controller,
	mdw middleware.Middleware,
) {
	app.Get("/", ctr.RenderMain)
	app.Get("/tag/:slug", ctr.RenderTag)
	app.Get("/tags", ctr.RenderTags)
	app.Get("/terms", ctr.RenderTerm)
	app.Get("/:dictionary", ctr.RenderDict)
	app.Get("/:dictionary/:word", ctr.RenderWord)
	app.Get("/monit", mdw.BasicAuth, monitor.New())
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
