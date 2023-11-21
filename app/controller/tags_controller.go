package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderTags(c *fiber.Ctx) error {
	ctr.Data.Title = "Теги"
	return c.Render("views/main", ctr.Response())
}

func (ctr *AppController) RenderTag(c *fiber.Ctx) error {
	slug := c.Params("slug")

	tag, err := ctr.TagRepository.GetBySlug(slug)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Render("views/errors/error", ctr.ErrorResponse())
	}

	dictionaries, err := ctr.DictRepository.GetByTag(slug)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Render("views/errors/error", ctr.ErrorResponse())
	}

	ctr.Data.Title = tag.Title + " словари и справочники"
	ctr.Data.Dictionaries = dictionaries

	return c.Render("views/tag", ctr.Response())
}
