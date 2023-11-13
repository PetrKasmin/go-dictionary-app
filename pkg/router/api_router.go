package router

import (
	"github.com/app-dictionary/app/controller"
	"github.com/app-dictionary/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(
	app *fiber.App,
	ctr controller.Controller,
	mdw middleware.Middleware,
) {
	api := app.Group("/api", limiter.New())
	api.Get("/search", ctr.Search)
	api.Get("/sitemap", mdw.BasicAuth, ctr.SitemapGenerate)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
