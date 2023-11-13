package controller

import (
	"github.com/app-dictionary/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderMain(c *fiber.Ctx) error {
	dictionaries, err := ctr.DictRepository.GetAll()
	if err != nil {
		return c.Render("views/errors/500", ctr.GetError(err))
	}

	ctr.Data.Title = "Словари, энциклопедии и справочники"
	ctr.Data.DictionariesByChunks = helpers.GetChunks(dictionaries, 2)

	return c.Render("views/main", ctr.GetResponse())
}
