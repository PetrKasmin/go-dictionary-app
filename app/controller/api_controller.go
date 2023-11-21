package controller

import (
	"fmt"
	"github.com/app-dictionary/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func (ctr *AppController) Search(c *fiber.Ctx) error {
	//word := c.Query("word")

	//words, err := ctr.WordRepository.Search(word)
	//log.Println("words", words)
	//if err != nil {
	//	return c.SendStatus(fiber.StatusInternalServerError)
	//}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"words": "words",
	})
}

func (ctr *AppController) SitemapGenerate(c *fiber.Ctx) error {
	dictionaries, err := ctr.DictRepository.GetAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var urls []string
	for i, d := range dictionaries {
		if i > 200 {
			//break
		}
		urls = append(urls, fmt.Sprintf("/%s", d.Slug))

		words, err := ctr.WordRepository.GetByDict(d.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		for _, w := range words {
			urls = append(urls, fmt.Sprintf("/%s/%s", d.Slug, w.Slug))
		}

		//fmt.Println("sitemap for dict "+d.Title, d.ID, len(urls))
	}

	helpers.SitemapGenerator(urls)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "generate ok",
	})
}
