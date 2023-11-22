package controller

import (
	"github.com/app-dictionary/app/helpers"
	"github.com/app-dictionary/app/models"
	"github.com/app-dictionary/pkg/redis"
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) RenderMain(c *fiber.Ctx) error {
	var letters []string
	var tags []models.Tag
	var dictionaries []models.Dictionary

	cacheManager := redis.NewCacheManager()

	if err := cacheManager.Get("tags", &tags); err != nil {
		tags, err := ctr.TagRepository.GetAll()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/errors/error", ctr.ErrorResponse())
		}

		if err := cacheManager.Set("tags", tags); err != nil {
			return err
		}
	}

	if err := cacheManager.Get("dictionaries", &dictionaries); err != nil {
		dictionaries, err = ctr.DictRepository.GetAll()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/errors/error", ctr.ErrorResponse())
		}

		dictionaries, letters = helpers.SplitByLetter(dictionaries)

		if err := cacheManager.Set("dictionaries", dictionaries); err != nil {
			return err
		}

		if err := cacheManager.Set("letters", letters); err != nil {
			return err
		}
	} else {
		if err := cacheManager.Get("letters", &letters); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.Render("views/errors/error", ctr.ErrorResponse())
		}
	}

	ctr.Data.Title = "Словари, энциклопедии и справочники"
	ctr.Data.Tags = tags
	ctr.Data.Dictionaries = dictionaries
	ctr.Data.DictionariesLetters = letters

	return c.Render("views/main", ctr.Response())
}
