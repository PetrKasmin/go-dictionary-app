package router

import (
	"github.com/app-dictionary/app/controller"
	"github.com/app-dictionary/app/middleware"
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	InstallRouter(app *fiber.App, ctr controller.Controller, mdw middleware.Middleware)
}
