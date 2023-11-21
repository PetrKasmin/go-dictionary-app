package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderMain(c *fiber.Ctx) error {
	tags, err := ctr.TagRepository.GetAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Render("views/errors/error", ctr.ErrorResponse())
	}

	dictionaries, err := ctr.DictRepository.GetAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Render("views/errors/error", ctr.ErrorResponse())
	}

	ctr.Data.Title = "Словари, энциклопедии и справочники"
	ctr.Data.Tags = tags
	ctr.Data.Dictionaries = dictionaries

	return c.Render("views/main", ctr.Response())
}
