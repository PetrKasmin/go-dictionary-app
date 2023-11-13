package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderTags(c *fiber.Ctx) error {
	ctr.Data.Title = "Теги"
	return c.Render("views/main", ctr.GetResponse())
}
