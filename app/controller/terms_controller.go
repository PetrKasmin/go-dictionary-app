package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderTerm(c *fiber.Ctx) error {
	ctr.Data.Title = "Пользовательское соглашение"
	return c.Render("views/terms", ctr.Response())
}
